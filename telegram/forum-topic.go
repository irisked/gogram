package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// CreateForumTopic performs createForumTopic telegram api method.
//
// options:
//
// option.CreateForumTopicIconColor(), option.CreateForumTopicIconEmojiID()
//
// Reference: https://core.telegram.org/bots/api#createforumtopic
func (t *Client) CreateForumTopic(ctx context.Context, chatID int64, name string, opts ...method.Option[*method.CreateForumTopic]) (method.ForumTopic, error) {
	m := method.NewCreateForumTopic(chatID, name, opts...)
	var result method.ForumTopic
	err := t.process(ctx, m, &result)
	return result, err
}

// EditForumTopic performs editForumTopic telegram api method.
//
// options:
//
// option.EditForumTopicName(), option.EditForumTopicIconEmojiID()
//
// Reference: https://core.telegram.org/bots/api#editforumtopic
func (t *Client) EditForumTopic(ctx context.Context, chatID int64, messageID int, opts ...method.Option[*method.EditForumTopic]) (bool, error) {
	m := method.NewEditForumTopic(chatID, messageID, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteForumTopic performs deleteForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deleteforumtopic
func (t *Client) DeleteForumTopic(ctx context.Context, chatID int64, messageID int) (bool, error) {
	m := method.NewDeleteForumTopic(chatID, messageID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetForumTopic performs getForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getforumtopic
func (t *Client) GetForumTopic(ctx context.Context, chatID int64, messageID int) (method.ForumTopic, error) {
	m := method.NewGetForumTopic(chatID, messageID)
	var result method.ForumTopic
	err := t.process(ctx, m, &result)
	return result, err
}

// CloseForumTopic performs closeForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#closeforumtopic
func (t *Client) CloseForumTopic(ctx context.Context, chatID int64, messageID int) (bool, error) {
	m := method.NewCloseForumTopic(chatID, messageID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// ReopenForumTopic performs reopenForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#reopenforumtopic
func (t *Client) ReopenForumTopic(ctx context.Context, chatID int64, messageID int) (bool, error) {
	m := method.NewReopenForumTopic(chatID, messageID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetForumTopicIconStickers performs getForumTopicIconStickers telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getforumtopiciconstickers
func (t *Client) GetForumTopicIconStickers(ctx context.Context) ([]types.Sticker, error) {
	m := method.NewGetForumTopicIconStickers()
	var result []types.Sticker
	err := t.process(ctx, m, &result)
	return result, err
}

// UnpinAllForumTopicMessages performs unpinAllForumTopicMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (t *Client) UnpinAllForumTopicMessages(ctx context.Context, chatID int64, messageID int) (bool, error) {
	m := method.NewUnpinAllForumTopicMessages(chatID, messageID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// EditGeneralTopic performs editGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#editgeneralforumtopic
func (t *Client) EditGeneralTopic(ctx context.Context, chatID int64, name string) (bool, error) {
	m := method.NewEditGeneralForumTopic(chatID, name)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// CloseGeneralForumTopic performs closeGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#closegeneralforumtopic
func (t *Client) CloseGeneralForumTopic(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewCloseGeneralForumTopic(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// ReopenGeneralForumTopic performs reopenGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#reopengeneralforumtopic
func (t *Client) ReopenGeneralForumTopic(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewReopenGeneralForumTopic(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// HideGeneralForumTopic performs hideGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#hidegeneralforumtopic
func (t *Client) HideGeneralForumTopic(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewHideGeneralForumTopic(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// UnhideGeneralForumTopic performs unhideGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (t *Client) UnhideGeneralForumTopic(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewUnhideGeneralForumTopic(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// UnpinAllGeneralForumTopicMessages performs unpinAllGeneralForumTopicMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (t *Client) UnpinAllGeneralForumTopicMessages(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewUnpinAllGeneralForumTopicMessages(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
