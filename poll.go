package gogram

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/irisked/gogram/telegram"
)

const (
	minBackoff = time.Second
	maxBackoff = 30 * time.Second
)

// StartPolling receives updates with long polling until ctx is
// cancelled, then waits for running handlers to return and returns nil.
//
// Each chat's updates are handled one at a time and in order, while
// different chats are handled concurrently; see WithConcurrency.
//
// Cancelling ctx stops polling but not handlers already running: they
// get a context that carries ctx's values without its cancellation, so
// their API calls can complete during shutdown. WithShutdownTimeout
// bounds that wait.
//
// Failed API calls are logged and retried with backoff. Polling
// doesn't work while a webhook is set; remove it with
// Client().DeleteWebhook first.
func (b *Bot) StartPolling(ctx context.Context) error {
	b.mu.Lock()
	if b.polling {
		b.mu.Unlock()
		return ErrAlreadyPolling
	}
	b.polling = true
	b.mu.Unlock()
	defer func() {
		b.mu.Lock()
		b.polling = false
		b.mu.Unlock()
	}()

	if !b.retry(ctx, "getMe", b.Init) {
		return nil
	}

	handlerCtx, cancelHandlers := context.WithCancel(context.WithoutCancel(ctx))
	defer cancelHandlers()
	d := newDispatcher(b.concurrency, func(u telegram.Update) { b.handle(handlerCtx, u) })

	offset := 0
	opts := []telegram.GetUpdatesOption{
		telegram.WithTimeout(int(b.pollTimeout / time.Second)),
		telegram.WithAllowedUpdates(b.allowedUpdates...),
	}
	if b.allowedUpdates == nil {
		opts = opts[:1] // keep Telegram's setting
	}
poll:
	for {
		var updates []telegram.Update
		ok := b.retry(ctx, "getUpdates", func(ctx context.Context) error {
			var err error
			updates, err = b.client.GetUpdates(ctx, append(opts, telegram.WithOffset(offset))...)
			return err
		})
		if !ok {
			break
		}
		for _, u := range updates {
			if !d.dispatch(ctx, u) {
				break poll
			}
			offset = u.UpdateID + 1
		}
	}

	err := b.drain(d, cancelHandlers)
	b.confirm(context.WithoutCancel(ctx), offset)
	return err
}

// retry calls fn until it succeeds, waiting with exponential backoff (or
// as long as Telegram asks) between attempts. It reports false if ctx
// ended first.
func (b *Bot) retry(ctx context.Context, what string, fn func(context.Context) error) bool {
	backoff := minBackoff
	for {
		err := fn(ctx)
		if ctx.Err() != nil {
			return false
		}
		if err == nil {
			return true
		}
		wait := backoff
		if apiErr, ok := errors.AsType[*telegram.Error](err); ok && apiErr.RetryAfter > 0 {
			wait = apiErr.RetryAfter
		}
		b.logger.WarnContext(ctx, "gogram: "+what+" failed", "error", err, "retry_in", wait)
		if !sleep(ctx, wait) {
			return false
		}
		backoff = min(backoff*2, maxBackoff)
	}
}

// drain waits for the dispatched updates to be handled, for at most the
// shutdown timeout.
func (b *Bot) drain(d *dispatcher, cancelHandlers context.CancelFunc) error {
	done := make(chan struct{})
	go func() {
		d.wait()
		close(done)
	}()
	if b.shutdownTimeout <= 0 {
		<-done
		return nil
	}
	t := time.NewTimer(b.shutdownTimeout)
	defer t.Stop()
	select {
	case <-done:
		return nil
	case <-t.C:
		cancelHandlers()
		b.logger.Warn("gogram: shutdown timeout expired; cancelled running handlers")
		return ErrShutdownTimeout
	}
}

// confirm tells Telegram that updates before offset were handled, so
// they aren't delivered again after a restart. Telegram otherwise only
// learns this from the next getUpdates call.
func (b *Bot) confirm(ctx context.Context, offset int) {
	if offset == 0 {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := b.client.GetUpdates(ctx, telegram.WithOffset(offset), telegram.WithLimit(1))
	if err != nil {
		b.logger.WarnContext(ctx, "gogram: confirming handled updates failed", "error", err)
	}
}

// sleep waits for d or until ctx is done, and reports whether the full
// duration elapsed.
func sleep(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-t.C:
		return true
	}
}

// dispatcher runs updates concurrently, except that updates with the
// same key (chat or user) run one at a time, in the order dispatched.
type dispatcher struct {
	handle func(telegram.Update)
	slots  chan struct{} // limits updates in flight; nil for no limit
	wg     sync.WaitGroup

	mu     sync.Mutex
	queues map[int64][]telegram.Update // a key is present while its worker runs
}

func newDispatcher(limit int, handle func(telegram.Update)) *dispatcher {
	d := &dispatcher{handle: handle, queues: map[int64][]telegram.Update{}}
	if limit > 0 {
		d.slots = make(chan struct{}, limit)
	}
	return d
}

// dispatch schedules u, waiting for a free slot if the limit is
// reached. It reports false if ctx ended first.
func (d *dispatcher) dispatch(ctx context.Context, u telegram.Update) bool {
	if d.slots != nil {
		select {
		case d.slots <- struct{}{}:
		case <-ctx.Done():
			return false
		}
	}
	d.wg.Add(1)
	key, ordered := orderKey(&u)
	if !ordered {
		go d.run(u)
		return true
	}
	d.mu.Lock()
	queue, running := d.queues[key]
	d.queues[key] = append(queue, u)
	d.mu.Unlock()
	if !running {
		go d.work(key)
	}
	return true
}

// work handles the queue of key until it is empty.
func (d *dispatcher) work(key int64) {
	for {
		d.mu.Lock()
		queue := d.queues[key]
		if len(queue) == 0 {
			delete(d.queues, key)
			d.mu.Unlock()
			return
		}
		u := queue[0]
		d.queues[key] = queue[1:]
		d.mu.Unlock()
		d.run(u)
	}
}

func (d *dispatcher) run(u telegram.Update) {
	defer func() {
		if d.slots != nil {
			<-d.slots
		}
		d.wg.Done()
	}()
	d.handle(u)
}

func (d *dispatcher) wait() { d.wg.Wait() }

// orderKey returns the key of the updates that must be handled in order
// with u: its chat's, or else its user's. A private chat has the ID of
// its user, so a user's messages and inline queries share a key.
func orderKey(raw *telegram.Update) (int64, bool) {
	u := NewUpdate(nil, *raw)
	if id := u.ChatID(); id != 0 {
		return id, true
	}
	if from := u.From(); from != nil {
		return from.ID, true
	}
	return 0, false
}
