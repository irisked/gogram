package gogram

import (
	"sync"
	"github.com/irisked/gogram/telegram"
	"github.com/irisked/gogram/types"
)

// Handler function.
type Handler func (ctx *Context)

// Gogram its a telegram bot framework
type Gogram struct {
	telegram *telegram.Telegram
	dispatcher Dispatcher
	poller poller
	exit chan interface{}
	updates chan types.Update 
}

// New creates Gogram
func New(token string) *Gogram {
	g := new(Gogram)
	g.telegram = telegram.New(token)
	g.poller = newPoller(g.telegram, 100)
	g.dispatcher = NewDispatcher()
	return g
}

func (g *Gogram) StartPolling() {
	g.startPolling()
}

func (g *Gogram) startPolling() {
	if !g.poller.isPolling() {
		var wg sync.WaitGroup
		g.updates = make(chan types.Update)
		g.exit = make(chan interface{})
		stop := g.poller.startPolling(g.updates)
		wg.Add(1)
		go func () {
			defer wg.Done()
			for {
				select {
				case <- g.exit:
					stop<-struct{}{}
					close(g.updates)
					return
				case update := <- g.updates:
					go g.handle(update)
				}
			}
		}()
		wg.Wait()
	} else {
		panic("*Gogram.startPolling() is already called.")
	}
}

func (g *Gogram) handle(update types.Update) {
	ctx := NewContext(*g.telegram, update)
	g.dispatcher.Handle(ctx)
}

func (g *Gogram) StopPolling() {
	g.stopPolling()
}

func (g *Gogram) stopPolling() {
	if g.poller.isPolling() {
		g.exit <- struct{}{}
	} else {
		panic("*Gogram.startPolling() is already called.")
	}
}

func(g *Gogram) On(ut types.UpdateType, handler Handler) {
	g.dispatcher.On(ut, handler)
}

func(g *Gogram) OnCallbackQuery(handler Handler) {
	g.dispatcher.OnCallbackQuery(handler)
}

func(g *Gogram) OnInlineQuery(handler Handler) {
	g.dispatcher.OnInlineQuery(handler)
}

func(g *Gogram) OnShippingQuery(handler Handler) {
	g.dispatcher.OnShippingQuery(handler)
}

func(g *Gogram) OnPreCheckoutQuery(handler Handler) {
	g.dispatcher.OnPreCheckoutQuery(handler)
}

func(g *Gogram) OnChosenInlineResult(handler Handler) {
	g.dispatcher.OnChosenInlineResult(handler)
}

func (g *Gogram) OnMessage(handler Handler) {
	g.dispatcher.OnMessage(handler)
}

func (g *Gogram) OnText(handler Handler) {
	g.dispatcher.OnText(handler)
}

func (g *Gogram) OnAudio(handler Handler) {
	g.dispatcher.OnAudio(handler)
}

func (g *Gogram) OnDocument(handler Handler) {
	g.dispatcher.OnDocument(handler)
}

func (g *Gogram) OnAnimation(handler Handler) {
	g.dispatcher.OnAnimation(handler)
}

func (g *Gogram) OnPhoto(handler Handler) {
	g.dispatcher.OnPhoto(handler)
}

func (g *Gogram) OnSticker(handler Handler) {
	g.dispatcher.OnSticker(handler)
}

func (g *Gogram) OnVideo(handler Handler) {
	g.dispatcher.OnVideo(handler)
}

func (g *Gogram) OnVideoNote(handler Handler) {
	g.dispatcher.OnVideoNote(handler)
}

func (g *Gogram) OnVoice(handler Handler) {
	g.dispatcher.OnVoice(handler)
}

func (g *Gogram) OnContact(handler Handler) {
	g.dispatcher.OnContact(handler)
}

func (g *Gogram) OnLocation(handler Handler) {
	g.dispatcher.OnLocation(handler)
}

func (g *Gogram) OnVenue(handler Handler) {
	g.dispatcher.OnVenue(handler)
}

func (g *Gogram) OnNewChatMembers(handler Handler) {
	g.dispatcher.OnNewChatMembers(handler)
}

func (g *Gogram) OnLeftChatMember(handler Handler) {
	g.dispatcher.OnLeftChatMember(handler)
}

func (g *Gogram) OnNewChatPhoto(handler Handler) {
	g.dispatcher.OnNewChatPhoto(handler)
}

func (g *Gogram) OnDeleteChatPhoto(handler Handler) {
	g.dispatcher.OnDeleteChatPhoto(handler)
}

func (g *Gogram) OnInvoice(handler Handler) {
	g.dispatcher.OnInvoice(handler)
}

func (g *Gogram) OnSuccessfulPayment(handler Handler) {
	g.dispatcher.OnSuccessfulPayment(handler)
}

func (g *Gogram) OnCommand(command string, handler Handler) {
	g.dispatcher.OnCommand(command, handler)
}

func (g *Gogram) OnStart(handler Handler) {
	g.dispatcher.OnStart(handler)
}

func (g *Gogram) OnHelp(handler Handler) {
	g.dispatcher.OnHelp(handler)
}
