package gogram

import (
	"context"
	"log"

	"github.com/irisked/gogram/telegram"
	"github.com/irisked/gogram/types"
)

// Handler function.
type Handler func(ctx *Context)

// Gogram its a telegram bot framework
type Gogram struct {
	telegram   *telegram.Client
	dispatcher *dispatcher
	poller     Poller
	cancel     context.CancelFunc
	updates    chan types.Update
}

// New creates Gogram
func New(token string) *Gogram {
	g := new(Gogram)
	g.telegram = telegram.New(token)
	g.poller = NewPoller(g.telegram, 100)
	g.dispatcher = newDispatcher()
	return g
}

// Errors returns transient errors encountered while polling. Reading
// from it is optional.
func (g *Gogram) Errors() <-chan error {
	return g.poller.Errors()
}

// StartPolling starts the bot using long polling. It returns
// immediately; polling stops when ctx is cancelled or StopPolling is
// called.
func (g *Gogram) StartPolling(ctx context.Context) error {
	if g.poller.IsPolling() {
		return ErrAlreadyPolling
	}
	ctx, cancel := context.WithCancel(ctx)
	g.cancel = cancel
	g.updates = make(chan types.Update)
	if err := g.poller.StartPolling(ctx, g.updates); err != nil {
		cancel()
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case update := <-g.updates:
				go g.handle(ctx, update)
			}
		}
	}()
	return nil
}

func (g *Gogram) handle(ctx context.Context, update types.Update) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("gogram: recovered panic in handler: %v", r)
		}
	}()
	c := NewContext(ctx, *g.telegram, update)
	g.dispatcher.Handle(c)
}

// StopPolling stops long polling.
func (g *Gogram) StopPolling() {
	if g.poller.IsPolling() && g.cancel != nil {
		g.cancel()
	}
}

// OnUpdate sets handler for update type.
func (g *Gogram) OnUpdate(ut types.UpdateType, handler Handler) {
	g.dispatcher.OnUpdate(ut, handler)
}

// OnPath sets handler for path.
func (g *Gogram) OnPath(path string, handler Handler) {
	g.dispatcher.OnPath(path, handler)
}

// OnCallbackQuery sets handler for callback query.
func (g *Gogram) OnCallbackQuery(handler Handler) {
	g.dispatcher.OnCallbackQuery(handler)
}

// OnCallback sets handler for specific callback query.
func (g *Gogram) OnCallback(path string, handler Handler) {
	g.dispatcher.OnCallback(path, handler)
}

// OnInlineQuery sets handler for inline query.
func (g *Gogram) OnInlineQuery(handler Handler) {
	g.dispatcher.OnInlineQuery(handler)
}

// OnShippingQuery sets handler for shipping query.
func (g *Gogram) OnShippingQuery(handler Handler) {
	g.dispatcher.OnShippingQuery(handler)
}

// OnPreCheckoutQuery sets handler for pre checkout query.
func (g *Gogram) OnPreCheckoutQuery(handler Handler) {
	g.dispatcher.OnPreCheckoutQuery(handler)
}

// OnChosenInlineResult sets handler for chosen inline result.
func (g *Gogram) OnChosenInlineResult(handler Handler) {
	g.dispatcher.OnChosenInlineResult(handler)
}

// OnMessage sets handler for message.
func (g *Gogram) OnMessage(handler Handler) {
	g.dispatcher.OnMessage(handler)
}

// OnEditedMessage sets handler for edited message.
func (g *Gogram) OnEditedMessage(handler Handler) {
	g.dispatcher.OnEditedMessage(handler)
}

// OnText sets handler for text message.
func (g *Gogram) OnText(handler Handler) {
	g.dispatcher.OnText(handler)
}

// OnAudio sets handler for audio message.
func (g *Gogram) OnAudio(handler Handler) {
	g.dispatcher.OnAudio(handler)
}

// OnDocument sets handler for document message.
func (g *Gogram) OnDocument(handler Handler) {
	g.dispatcher.OnDocument(handler)
}

// OnAnimation sets handler for animation message.
func (g *Gogram) OnAnimation(handler Handler) {
	g.dispatcher.OnAnimation(handler)
}

// OnPhoto sets handler for photo message.
func (g *Gogram) OnPhoto(handler Handler) {
	g.dispatcher.OnPhoto(handler)
}

// OnSticker sets handler for sticker message.
func (g *Gogram) OnSticker(handler Handler) {
	g.dispatcher.OnSticker(handler)
}

// OnVideo sets handler for video message.
func (g *Gogram) OnVideo(handler Handler) {
	g.dispatcher.OnVideo(handler)
}

// OnVideoNote sets handler for video note message.
func (g *Gogram) OnVideoNote(handler Handler) {
	g.dispatcher.OnVideoNote(handler)
}

// OnVoice sets handler for voice message.
func (g *Gogram) OnVoice(handler Handler) {
	g.dispatcher.OnVoice(handler)
}

// OnContact sets handler for contact message.
func (g *Gogram) OnContact(handler Handler) {
	g.dispatcher.OnContact(handler)
}

// OnLocation sets handler for location message.
func (g *Gogram) OnLocation(handler Handler) {
	g.dispatcher.OnLocation(handler)
}

// OnVenue sets handler for venue message.
func (g *Gogram) OnVenue(handler Handler) {
	g.dispatcher.OnVenue(handler)
}

// OnNewChatMembers sets handler for new chat members.
func (g *Gogram) OnNewChatMembers(handler Handler) {
	g.dispatcher.OnNewChatMembers(handler)
}

// OnLeftChatMember sets handler for left chat member.
func (g *Gogram) OnLeftChatMember(handler Handler) {
	g.dispatcher.OnLeftChatMember(handler)
}

// OnNewChatPhoto sets handler for new chat photo.
func (g *Gogram) OnNewChatPhoto(handler Handler) {
	g.dispatcher.OnNewChatPhoto(handler)
}

// OnDeleteChatPhoto sets handler for delete chat photo.
func (g *Gogram) OnDeleteChatPhoto(handler Handler) {
	g.dispatcher.OnDeleteChatPhoto(handler)
}

// OnInvoice sets handler for invoice.
func (g *Gogram) OnInvoice(handler Handler) {
	g.dispatcher.OnInvoice(handler)
}

// OnSuccessfulPayment sets handler for successfull payment.
func (g *Gogram) OnSuccessfulPayment(handler Handler) {
	g.dispatcher.OnSuccessfulPayment(handler)
}

// OnCommand sets handler for command message.
func (g *Gogram) OnCommand(command string, handler Handler) {
	g.dispatcher.OnCommand(command, handler)
}

// OnStart sets handler for start command.
func (g *Gogram) OnStart(handler Handler) {
	g.dispatcher.OnStart(handler)
}

// OnHelp sets handler for help command.
func (g *Gogram) OnHelp(handler Handler) {
	g.dispatcher.OnHelp(handler)
}
