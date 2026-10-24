package telegram

import (
	"context"

	"github.com/irisked/gogram/content"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// CopyMessage performs copyMessage telegram api method.
//
// options:
//
// option.CopyMessageCaption(), option.CopyMessageParseMode(),
// option.CopyMessageDisableNotification(), option.CopyMessageReplyToMessageID(id int),
// option.CopyMessageAllowCopy(), option.CopyMessageReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#copymessage
func (t *Client) CopyMessage(ctx context.Context, chatID, fromChatID int64, messageID int, opts ...method.Option[*method.CopyMessage]) (types.MessageID, error) {
	m := method.NewCopyMessage(chatID, fromChatID, messageID, opts...)
	var result types.MessageID
	err := t.process(ctx, m, &result)
	return result, err
}

// CopyMessages performs copyMessages telegram api method.
//
// options:
//
// option.CopyMessagesDisableNotification(), option.CopyMessagesRemoveSender()
//
// Reference: https://core.telegram.org/bots/api#copymessages
func (t *Client) CopyMessages(ctx context.Context, chatID, fromChatID int64, messageIDs []int, opts ...method.Option[*method.CopyMessages]) ([]types.MessageID, error) {
	m := method.NewCopyMessages(chatID, fromChatID, messageIDs, opts...)
	var result []types.MessageID
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteMessages performs deleteMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletemessages
func (t *Client) DeleteMessages(ctx context.Context, chatID int64, messageIDs []int) (bool, error) {
	m := method.NewDeleteMessages(chatID, messageIDs)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// StopPoll performs stopPoll telegram api method.
//
// options:
//
// option.StopPollReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#stoppoll
func (t *Client) StopPoll(ctx context.Context, chatID int64, messageID int, opts ...method.Option[*method.StopPoll]) (types.Poll, error) {
	m := method.NewStopPoll(chatID, messageID, opts...)
	var result types.Poll
	err := t.process(ctx, m, &result)
	return result, err
}

// EditMessageLiveLocation performs editMessageLiveLocation telegram api method.
//
// options:
//
// option.EditMessageLiveLocationReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#editmessagelivelocation
func (t *Client) EditMessageLiveLocation(ctx context.Context, messageID *content.MessageTarget, location types.Location, opts ...method.Option[*method.EditMessageLiveLocation]) (*types.Message, error) {
	m := method.NewEditMessageLiveLocation(messageID, location, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// StopMessageLiveLocation performs stopMessageLiveLocation telegram api method.
//
// options:
//
// option.StopMessageLiveLocationReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#stopmessagelivelocation
func (t *Client) StopMessageLiveLocation(ctx context.Context, messageID *content.MessageTarget, opts ...method.Option[*method.StopMessageLiveLocation]) (*types.Message, error) {
	m := method.NewStopMessageLiveLocation(messageID, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// DeleteWebhook performs deleteWebhook telegram api method.
//
// options:
//
// option.DeleteWebhookDropPendingUpdates()
//
// Reference: https://core.telegram.org/bots/api#deletewebhook
func (t *Client) DeleteWebhook(ctx context.Context, opts ...method.Option[*method.DeleteWebhook]) (bool, error) {
	m := method.NewDeleteWebhook(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
