package telegram

import (
	"context"

	"github.com/irisked/gogram/content"
	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SendMessage performs sendMessage telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.DisableNotification()
// option.ReplyToMessageID(id int), option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendmessage
func (t *Client) SendMessage(ctx context.Context, chatID int64, text content.TextData, opts ...method.Option[*method.SendMessage]) (types.Message, error) {
	m := method.NewSendMessage(chatID, text, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// ForwardMessage performs forwardMessage telegram api method.
//
// options:
//
// option.DisableNotification()
//
// Reference: https://core.telegram.org/bots/api#forwardmessage
func (t *Client) ForwardMessage(ctx context.Context, chatID, fromChatID, messageID int, opts ...method.Option[*method.ForwardMessage]) (types.Message, error) {
	m := method.NewForwardMessage(chatID, fromChatID, messageID, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// EditMessageText performs editMessageText telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.DisableWebPagePreview(), option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagetext
func (t *Client) EditMessageText(ctx context.Context, messageID *content.MessageID, text content.TextData, opts ...method.Option[*method.EditMessageText]) (*types.Message, error) {
	m := method.NewEditMessageText(messageID, text, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// EditMessageCaption performs editMessageCaption telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagecaption
func (t *Client) EditMessageCaption(ctx context.Context, messageID *content.MessageID, caption content.TextData, opts ...method.Option[*method.EditMessageCaption]) (*types.Message, error) {
	m := method.NewEditMessageCaption(messageID, caption, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// EditMessageReplyMarkup performs editMessageReplyMarkup telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// Reference: https://core.telegram.org/bots/api#editmessagereplymarkup
func (t *Client) EditMessageReplyMarkup(ctx context.Context, messageID *content.MessageID, kb keyboard.Markup) (*types.Message, error) {
	m := method.NewEditMessageReplyMarkup(messageID, kb)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// DeleteMessage performs deleteMessage telegram api method.
//
// https://core.telegram.org/bots/api#deletemessage
func (t *Client) DeleteMessage(ctx context.Context, chatID int64, messageID int) (bool, error) {
	m := method.NewDeleteMessage(chatID, messageID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
