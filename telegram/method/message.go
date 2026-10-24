package method

import (
	"github.com/irisked/gogram/args"
	"github.com/irisked/gogram/telegram/keyboard"
)

// Option is a generic functional option type
type Option[T any] func(T)

// EditOptions contains fields common to edit methods
type EditOptions struct {
	ChatID          int64  `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty" form:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty" form:"inline_message_id,omitempty"`
}

// SendMessage contains information about sendMessage telegram api method.
type SendMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
	SendOptions
	SimpleMethod
}

// NewSendMessage creates SendMessage.
func NewSendMessage(chatID int64, text args.TextData, opts ...Option[*SendMessage]) *SendMessage {
	m := &SendMessage{
		ChatID: chatID,
		Text:   text.Text,
	}
	m.ParseMode = text.ParseMode
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SendMessage) Method() string {
	return "sendMessage"
}

// EditMessageText contains information about editMessageText telegram api method.
type EditMessageText struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Text            string `json:"text"`
	EditOptions
	SimpleMethod
}

// NewEditMessageText creates EditMessageText.
func NewEditMessageText(message *args.MessageID, text args.TextData, opts ...Option[*EditMessageText]) *EditMessageText {
	m := &EditMessageText{
		ChatID:          message.ChatID,
		MessageID:       message.MessageID,
		InlineMessageID: message.InlineMessageID,
		Text:            text.Text,
	}
	for _, o := range opts {
		o(m)
	}
	return m
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

// NewForwardMessage creates ForwardMessage.
func NewForwardMessage(chatID, fromChatID, messageID int, opts ...Option[*ForwardMessage]) *ForwardMessage {
	m := &ForwardMessage{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *ForwardMessage) Method() string {
	return "forwardMessage"
}

// EditMessageCaption contains information about editMessageCaption telegram api method.
type EditMessageCaption struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	CaptionOptions
	EditOptions
	SimpleMethod
}

// NewEditMessageCaption creates EditMessageCaption
func NewEditMessageCaption(message *args.MessageID, text args.TextData, opts ...Option[*EditMessageCaption]) *EditMessageCaption {
	m := &EditMessageCaption{
		ChatID:          message.ChatID,
		MessageID:       message.MessageID,
		InlineMessageID: message.InlineMessageID,
	}
	m.Caption = text.Text
	m.ParseMode = text.ParseMode
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *EditMessageCaption) Method() string {
	return "editMessageCaption"
}

// EditMessageReplyMarkup contains information about editMessageReplyMarkup telegram api method.
type EditMessageReplyMarkup struct {
	ChatID          int64           `json:"chat_id,omitempty"`
	MessageID       int             `json:"message_id,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty"`
	ReplyMarkup     keyboard.Markup `json:"reply_markup,omitempty"`
	EditOptions
	SimpleMethod
}

// Method returns telegram api method endpoint.
func (r *EditMessageReplyMarkup) Method() string {
	return "editMessageReplyMarkup"
}

// NewEditMessageReplyMarkup creates EditMessageReplyMarkup.
func NewEditMessageReplyMarkup(message *args.MessageID, kb keyboard.Markup, opts ...Option[*EditMessageReplyMarkup]) *EditMessageReplyMarkup {
	m := &EditMessageReplyMarkup{
		ChatID:          message.ChatID,
		MessageID:       message.MessageID,
		InlineMessageID: message.InlineMessageID,
		ReplyMarkup:     kb,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// DeleteMessage contains information about deleteMessage telegram api method.
type DeleteMessage struct {
	ChatID    int64 `json:"chat_id,omitempty"`
	MessageID int   `json:"message_id,omitempty"`
	SimpleMethod
}

// NewDeleteMessage creates SendMessage
func NewDeleteMessage(chatID int64, messageID int) *DeleteMessage {
	m := &DeleteMessage{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteMessage) Method() string {
	return "deleteMessage"
}
