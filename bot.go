package gogram

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/irisked/gogram/internal/router"
	"github.com/irisked/gogram/telegram"
)

// Errors returned by StartPolling.
var (
	ErrAlreadyPolling  = errors.New("gogram: bot is already polling")
	ErrShutdownTimeout = errors.New("gogram: handlers still running after the shutdown timeout")
)

// Handler handles one update. A returned error is passed to the bot's
// ErrorHandler.
type Handler func(ctx context.Context, u *Update) error

// Middleware wraps a Handler, e.g. to log, authorize or time updates.
type Middleware func(next Handler) Handler

// ErrorHandler handles the errors returned by handlers, including
// recovered panics as a *PanicError. The default logs them with the
// bot's logger.
type ErrorHandler func(ctx context.Context, u *Update, err error)

// PanicError is the error passed to the ErrorHandler when a handler
// panics.
type PanicError struct {
	Value any    // the value passed to panic
	Stack []byte // the stack trace of the panicking goroutine
}

func (e *PanicError) Error() string { return fmt.Sprintf("gogram: handler panicked: %v", e.Value) }

// Unwrap returns the panic value if it is an error.
func (e *PanicError) Unwrap() error {
	err, _ := e.Value.(error)
	return err
}

// Bot is a Telegram bot. Register handlers before calling StartPolling
// or serving WebhookHandler; registering while updates are being
// handled is not safe.
type Bot struct {
	client          *telegram.Client
	logger          *slog.Logger
	onError         ErrorHandler
	pollTimeout     time.Duration
	allowedUpdates  []telegram.UpdateType
	concurrency     int
	shutdownTimeout time.Duration

	routes     *router.Tree[Handler]
	middleware []Middleware

	username atomic.Pointer[string]
	initMu   sync.Mutex

	mu      sync.Mutex
	polling bool
}

// Option configures a Bot created by New or NewWithClient.
type Option func(*Bot)

// WithLogger sets the logger for polling errors and, with the default
// ErrorHandler, handler errors. The default is slog.Default().
func WithLogger(l *slog.Logger) Option {
	return func(b *Bot) { b.logger = l }
}

// WithErrorHandler sets the function that handles errors returned by
// handlers and recovered panics.
func WithErrorHandler(h ErrorHandler) Option {
	return func(b *Bot) { b.onError = h }
}

// WithPollTimeout sets the long-polling timeout of getUpdates. The
// default is 30 seconds.
func WithPollTimeout(d time.Duration) Option {
	return func(b *Bot) { b.pollTimeout = d }
}

// WithAllowedUpdates sets the update types Telegram delivers while
// polling; telegram.AllUpdateTypes lists them all. Without it, Telegram
// keeps the setting of the last getUpdates or setWebhook call, which
// initially is every type except chat_member, message_reaction and
// message_reaction_count.
func WithAllowedUpdates(types ...telegram.UpdateType) Option {
	return func(b *Bot) { b.allowedUpdates = types }
}

// WithUsername sets the bot's username, without "@". The bot uses it to
// ignore commands addressed to other bots, such as "/start@other_bot"
// in a group. Without it, StartPolling and WebhookHandler look it up with
// Init.
func WithUsername(username string) Option {
	return func(b *Bot) { b.setUsername(username) }
}

// WithConcurrency limits how many updates StartPolling handles at once,
// counting the ones waiting behind an earlier update from the same
// chat. When the limit is reached, polling waits. The default, 0, means
// no limit.
//
// Updates from the same chat (or, for updates without a chat, from the
// same user) are always handled one at a time, in order.
func WithConcurrency(n int) Option {
	return func(b *Bot) { b.concurrency = n }
}

// WithShutdownTimeout limits how long StartPolling waits for running
// handlers once its context is cancelled. When it expires, the handlers'
// context is cancelled too and StartPolling returns ErrShutdownTimeout.
// The default, 0, waits for as long as handlers take.
func WithShutdownTimeout(d time.Duration) Option {
	return func(b *Bot) { b.shutdownTimeout = d }
}

// New creates a Bot for the given token.
func New(token string, opts ...Option) *Bot {
	return NewWithClient(telegram.New(token), opts...)
}

// NewWithClient creates a Bot that uses an existing client, e.g. one
// configured with telegram.WithBaseURL.
func NewWithClient(client *telegram.Client, opts ...Option) *Bot {
	b := &Bot{
		client:      client,
		pollTimeout: 30 * time.Second,
		routes:      router.New[Handler](),
	}
	for _, o := range opts {
		o(b)
	}
	if b.logger == nil {
		b.logger = slog.Default()
	}
	if b.onError == nil {
		b.onError = b.logError
	}
	return b
}

// Client returns the Telegram Bot API client.
func (b *Bot) Client() *telegram.Client {
	return b.client
}

// Use adds middleware that runs around every handler, in the order
// added.
func (b *Bot) Use(mw ...Middleware) {
	b.middleware = append(b.middleware, mw...)
}

// Init looks up the bot's username with getMe; see WithUsername.
// StartPolling and WebhookHandler call it when needed, so call it
// yourself only to fail early on a bad token, or before using
// HandleUpdate directly.
func (b *Bot) Init(ctx context.Context) error {
	b.initMu.Lock()
	defer b.initMu.Unlock()
	if b.Username() != "" {
		return nil
	}
	me, err := b.client.GetMe(ctx)
	if err != nil {
		return err
	}
	b.setUsername(me.Username)
	return nil
}

// Username returns the bot's username, or "" until it is known.
func (b *Bot) Username() string {
	if p := b.username.Load(); p != nil {
		return *p
	}
	return ""
}

func (b *Bot) setUsername(name string) {
	name = strings.TrimPrefix(name, "@")
	b.username.Store(&name)
}

// HandleUpdate routes a single update to its handler, waits for it to
// return and returns its error. A panic in the handler is recovered and
// returned as a *PanicError. HandleUpdate doesn't call the ErrorHandler:
// StartPolling and WebhookHandler do.
//
// Call it directly to feed updates from another source, or in tests.
func (b *Bot) HandleUpdate(ctx context.Context, update telegram.Update) (err error) {
	_, h, ok := b.routes.Match(routeOf(&update, b.Username()))
	if !ok {
		return nil
	}
	for i := len(b.middleware) - 1; i >= 0; i-- {
		h = b.middleware[i](h)
	}

	defer func() {
		if r := recover(); r != nil {
			err = &PanicError{Value: r, Stack: debug.Stack()}
		}
	}()
	return h(ctx, NewUpdate(b.client, update))
}

// handle runs HandleUpdate and reports its error.
func (b *Bot) handle(ctx context.Context, update telegram.Update) {
	if err := b.HandleUpdate(ctx, update); err != nil {
		b.onError(ctx, NewUpdate(b.client, update), err)
	}
}

func (b *Bot) logError(ctx context.Context, u *Update, err error) {
	attrs := []any{"update_id", u.Raw().UpdateID, "error", err}
	if p, ok := errors.AsType[*PanicError](err); ok {
		attrs = append(attrs, "stack", string(p.Stack))
	}
	b.logger.ErrorContext(ctx, "gogram: handler failed", attrs...)
}
