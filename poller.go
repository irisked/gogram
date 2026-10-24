package gogram

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/irisked/gogram/telegram"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// ErrAlreadyPolling is returned by StartPolling when the poller is
// already running.
var ErrAlreadyPolling = errors.New("gogram: poller is already running")

// Poller interface.
type Poller interface {
	StartPolling(ctx context.Context, updates chan<- types.Update) error
	IsPolling() bool
	// Errors returns transient errors encountered while polling (e.g.
	// a failed GetUpdates call). Reading from it is optional; errors
	// that aren't read are dropped rather than blocking the poller.
	Errors() <-chan error
}

type pollingConfig struct {
	mu      sync.Mutex
	offset  int
	limit   int
	started bool
}

func (c *pollingConfig) isStarted() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.started
}

func (c *pollingConfig) setStarted(v bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.started = v
}

func (c *pollingConfig) offsetAndLimit() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.offset, c.limit
}

func (c *pollingConfig) setOffset(v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.offset = v
}

type poller struct {
	config pollingConfig
	tg     *telegram.Client
	errs   chan error
}

// NewPoller creates poller.
func NewPoller(tg *telegram.Client, limit int) Poller {
	p := poller{
		config: pollingConfig{
			offset: -1,
			limit:  limit,
		},
		tg:   tg,
		errs: make(chan error, 16),
	}
	return &p
}

func (p *poller) IsPolling() bool {
	return p.config.isStarted()
}

func (p *poller) Errors() <-chan error {
	return p.errs
}

func (p *poller) StartPolling(ctx context.Context, updates chan<- types.Update) error {
	if p.IsPolling() {
		return ErrAlreadyPolling
	}
	p.config.setStarted(true)
	go func() {
		defer p.config.setStarted(false)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if err := p.fetch(ctx, updates); err != nil {
				p.reportErr(err)
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
			}
		}
	}()
	return nil
}

func (p *poller) reportErr(err error) {
	select {
	case p.errs <- err:
	default:
	}
}

func (p *poller) fetch(ctx context.Context, updates chan<- types.Update) error {
	offset, limit := p.config.offsetAndLimit()
	incoming, err := p.tg.GetUpdates(ctx,
		func(u *method.GetUpdates) { u.Offset = offset },
		func(u *method.GetUpdates) { u.Limit = limit },
	)
	if err != nil {
		return err
	}
	if len(incoming) > 0 {
		p.config.setOffset(incoming[len(incoming)-1].UpdateID + 1)
		for _, update := range incoming {
			updates <- update
		}
	}
	return nil
}
