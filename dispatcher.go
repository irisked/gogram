package gogram

import (
	"github.com/irisked/gogram/types"
	"fmt"
	"strings"
)


// Dispatcher interface
type Dispatcher interface {
	OnPath(string, Handler)
	OnUpdate(types.UpdateType, Handler)
	OnCallbackQuery(Handler)
	OnCallback(string, Handler)
	OnInlineQuery(Handler)
	OnShippingQuery(Handler)
	OnPreCheckoutQuery(Handler)
	OnChosenInlineResult(Handler)
	OnMessage(Handler)
	OnEditedMessage(Handler)
	OnText(Handler)
	OnAudio(Handler)
	OnDocument(Handler)
	OnAnimation(Handler)
	OnPhoto(Handler)
	OnSticker(Handler)
	OnVideo(Handler)
	OnVideoNote(Handler)
	OnVoice(Handler)
	OnContact(Handler)
	OnLocation(Handler)
	OnVenue(Handler)
	OnNewChatMembers(Handler)
	OnLeftChatMember(Handler)
	OnNewChatPhoto(Handler)
	OnDeleteChatPhoto(Handler)
	OnInvoice(Handler)
	OnSuccessfulPayment(Handler)
	OnCommand(string, Handler)
	OnStart(Handler)
	OnHelp(Handler)
	Handle(*Context)
}

type dispatcher struct {
	storage Storage
}

// NewDispatcher creates dispatcher
func NewDispatcher() Dispatcher {
	c := new(dispatcher)
	c.storage = NewStorage()
	return c
}

func(c *dispatcher) OnUpdate(ut types.UpdateType, handler Handler) {
	c.OnPath(defaultPath(ut), handler)
}

func(c *dispatcher) OnPath(path string, handler Handler) {
	c.storage.Insert(path, handler)
}

func(c *dispatcher) OnCallbackQuery(handler Handler) {
	c.OnPath(defaultPath(types.IncomingCallbackQuery), handler)
}

func(c *dispatcher) OnCallback(path string, handler Handler) {
	c.OnPath(defaultStringPath(types.IncomingCallbackQuery.String(), path), handler)
}

func(c *dispatcher) OnInlineQuery(handler Handler) {
	c.OnPath(defaultPath(types.IncomingInlineQuery), handler)
}

func(c *dispatcher) OnShippingQuery(handler Handler) {
	c.OnPath(defaultPath(types.IncomingShippingQuery), handler)
}

func(c *dispatcher) OnPreCheckoutQuery(handler Handler) {
	c.OnPath(defaultPath(types.IncomingPreCheckoutQuery), handler)
}

func(c *dispatcher) OnChosenInlineResult(handler Handler) {
	c.OnPath(defaultPath(types.IncomingChosenInlineResult), handler)
}

func (c *dispatcher) OnMessage(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage), handler)
}

func (c *dispatcher) OnEditedMessage(handler Handler) {
	c.OnPath(defaultPath(types.IncomingEditedMessage), handler)
}

func (c *dispatcher) OnText(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.TextMessage), handler)
}

func (c *dispatcher) OnAudio(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.AudioMessage), handler)
}

func (c *dispatcher) OnDocument(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.DocumentMessage), handler)
}

func (c *dispatcher) OnAnimation(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.AnimationMessage), handler)
}

func (c *dispatcher) OnPhoto(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.PhotoMessage), handler)
}

func (c *dispatcher) OnSticker(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.StickerMessage), handler)
}

func (c *dispatcher) OnVideo(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.VideoMessage), handler)
}

func (c *dispatcher) OnVideoNote(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.VideoNoteMessage), handler)
}

func (c *dispatcher) OnVoice(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.VoiceMessage), handler)
}

func (c *dispatcher) OnContact(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.ContactMessage), handler)
}

func (c *dispatcher) OnLocation(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.LocationMessage), handler)
}

func (c *dispatcher) OnVenue(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.VenueMessage), handler)
}

func (c *dispatcher) OnNewChatMembers(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.NewChatMembersMessage), handler)
}

func (c *dispatcher) OnLeftChatMember(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.LeftChatMemberMessage), handler)
}

func (c *dispatcher) OnNewChatPhoto(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.NewChatPhotoMessage), handler)
}

func (c *dispatcher) OnDeleteChatPhoto(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.DeleteChatPhotoMessage), handler)
}

func (c *dispatcher) OnInvoice(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.InvoiceMessage), handler)
}

func (c *dispatcher) OnSuccessfulPayment(handler Handler) {
	c.OnPath(defaultPath(types.IncomingMessage, types.SuccessfulPaymentMessage), handler)
}

func (c *dispatcher) OnCommand(command string, handler Handler) {
	c.OnPath(defaultStringPath(types.IncomingMessage.String(), types.CommandMessage.String(), command), handler)
}

func (c *dispatcher) OnStart(handler Handler) {
	c.OnCommand("start", handler)
}

func (c *dispatcher) OnHelp(handler Handler) {
	c.OnCommand("help", handler)
}

func (c *dispatcher) Handle(ctx *Context) {
	path := ctx.update.Path()
	if handler, ok := c.storage.Find(path); ok {
		handler(ctx)
	}
}

func defaultPath(elements ...fmt.Stringer) string {
	return path("/", elements)
}

func path(separator string, elements []fmt.Stringer) string {
	var builder strings.Builder
	for _, element := range elements {
		builder.WriteString(separator)
		builder.WriteString(element.String())
	}
	return builder.String()
}

func stringPath(separator string, elements []string) string {
	var builder strings.Builder
	for _, element := range elements {
		builder.WriteString(separator)
		builder.WriteString(element)
	}
	return builder.String()
}

func defaultStringPath(elements ...string) string {
	return stringPath("/", elements)
}