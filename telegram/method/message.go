package method

import (
	"github.com/irisked/gogram/params"
	"github.com/irisked/gogram/types/markup/keyboard"
)

// SendMessage contains information about sendMessage telegram api method.
type SendMessage struct {
	ChatID                int64             `json:"chat_id"`
	Text                  string            `json:"text"`
	ParseMode             string            `json:"parse_mode,omitempty" `
	DisableWebPagePreview bool              `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID      int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// SendMessageOption its a functional options interface.
type SendMessageOption interface {
	ApplySendMessageOption(*SendMessage)
}

// NewSendMessage creates SendMessage.
func NewSendMessage(chatID int64, text *params.Text, options ...SendMessageOption) *SendMessage {
	method := new(SendMessage)
	method.ChatID = chatID
	method.Text = text.Text
	method.ParseMode = text.ParseMode
	for _, option := range options {
		option.ApplySendMessageOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendMessage) Method() string {
	return "sendMessage"
}

// EditMessageText contains information about editMessageText telegram api method.
type EditMessageText struct {
	ChatID                int64             `json:"chat_id,omitempty"`
	MessageID             int               `json:"message_id,omitempty"`
	InlineMessageID       string            `json:"inline_message_id,omitempty"`
	Text                  string            `json:"text"`
	ParseMode             string            `json:"parse_mode,omitempty" `
	DisableWebPagePreview bool              `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// EditMessageTextOption its a functional options interface.
type EditMessageTextOption interface {
	ApplyEditMessageTextOption(*EditMessageText)
}

// NewEditMessageText creates EditMessageText.
func NewEditMessageText(message *params.MessageID, text params.Text, options ...EditMessageTextOption) *EditMessageText {
	method := new(EditMessageText)
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	method.Text = text.Text
	method.ParseMode = text.ParseMode
	for _, option := range options {
		option.ApplyEditMessageTextOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *EditMessageText) Method() string {
	return "editMessageText"
}

// ForwardMessage contains information about forwardMessage telegram api method.
type ForwardMessage struct {
	ChatID              int  `json:"chat_id"`
	FromChatID          int  `json:"from_chat_id"`
	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	SimpleMethod
}

// ForwardMessageOption its a functional options interface.
type ForwardMessageOption interface {
	ApplyForwardMessageOption(*ForwardMessage)
}

// NewForwardMessage creates ForwardMessage.
func NewForwardMessage(chatID, fromChatID, messageID int, options ...ForwardMessageOption) *ForwardMessage {
	method := new(ForwardMessage)
	method.ChatID = chatID
	method.FromChatID = fromChatID
	method.MessageID = messageID
	for _, option := range options {
		option.ApplyForwardMessageOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *ForwardMessage) Method() string {
	return "forwardMessage"
}

// EditMessageCaption contains information about editMessageCaption telegram api method.
type EditMessageCaption struct {
	ChatID          int64             `json:"chat_id,omitempty"`
	MessageID       int               `json:"message_id,omitempty"`
	InlineMessageID string            `json:"inline_message_id,omitempty"`
	Caption         string            `json:"caption"`
	ParseMode       string            `json:"parse_mode,omitempty"`
	ReplyMarkup     keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// Method returns telegram api method endpoint.
func (r *EditMessageCaption) Method() string {
	return "editMessageCaption"
}

// NewEditMessageCaption creates EditMessageCaption
func NewEditMessageCaption(message *params.MessageID, text params.Text, options ...EditMessageCaptionOption) *EditMessageCaption {
	method := new(EditMessageCaption)
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	method.Caption = text.Text
	method.ParseMode = text.ParseMode
	for _, option := range options {
		option.ApplyEditMessageCaptionOption(method)
	}
	return method
}

// EditMessageCaptionOption its a functional options interface.
type EditMessageCaptionOption interface {
	ApplyEditMessageCaptionOption(*EditMessageCaption)
}

// EditMessageReplyMarkup contains information about editMessageReplyMarkup telegram api method.
type EditMessageReplyMarkup struct {
	ChatID          int64             `json:"chat_id,omitempty"`
	MessageID       int               `json:"message_id,omitempty"`
	InlineMessageID string            `json:"inline_message_id,omitempty"`
	ReplyMarkup     keyboard.Keyboard `json:"reply_markup"`
	SimpleMethod
}

// Method returns telegram api method endpoint.
func (r *EditMessageReplyMarkup) Method() string {
	return "editMessageReplyMarkup"
}

// NewEditMessageReplyMarkup creates EditMessageReplyMarkup.
func NewEditMessageReplyMarkup(message *params.MessageID, keyboard keyboard.Keyboard) *EditMessageReplyMarkup {
	method := new(EditMessageReplyMarkup)
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	method.ReplyMarkup = keyboard
	return method
}

// DeleteMessage contains information about deleteMessage telegram api method.
type DeleteMessage struct {
	ChatID    int64 `json:"chat_id,omitempty"`
	MessageID int   `json:"message_id,omitempty"`
	SimpleMethod
}

// NewDeleteMessage creates SendMessage
func NewDeleteMessage(chatID int64, messageID int) *DeleteMessage {
	method := new(DeleteMessage)
	method.ChatID = chatID
	method.MessageID = messageID
	return method
}

// Method returns telegram api method endpoint.
func (r *DeleteMessage) Method() string {
	return "deleteMessage"
}
