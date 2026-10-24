package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// GetUserProfilePhotos performs getUserProfilePhotos telegram api method.
//
// options:
//
// option.Offset(o int), option.Limit(l int)
//
// Reference: https://core.telegram.org/bots/api#getuserprofilephotos
func (t *Client) GetUserProfilePhotos(ctx context.Context, userID int, opts ...method.Option[*method.GetUserProfilePhotos]) (types.UserProfilePhotos, error) {
	m := method.NewGetUserProfilePhotos(userID, opts...)
	photos := types.UserProfilePhotos{}
	err := t.process(ctx, m, &photos)
	return photos, err
}

// KickChatMember performs kickChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#kickchatmember
func (t *Client) KickChatMember(ctx context.Context, userID int, chatID int64, opts ...method.Option[*method.KickChatMember]) (bool, error) {
	m := method.NewKickChatMember(userID, chatID, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// UnbanChatMember performs unbanChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#unbanchatmember
func (t *Client) UnbanChatMember(ctx context.Context, userID int, chatID int64) (bool, error) {
	m := method.NewUnbanChatMember(userID, chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// RestrictChatMember performs restrictChatMember telegram api method.
//
// options:
//
// option.CanSendMessages(), options.CanSendMediaMessages(),
// option.CanSendOtherMessages(), option.CanAddWebPagePreview()
//
// Reference: https://core.telegram.org/bots/api#restrictchatmember
// TODO: rework true/false
func (t *Client) RestrictChatMember(ctx context.Context, userID int, chatID int64, opts ...method.Option[*method.RestrictChatMember]) (bool, error) {
	m := method.NewRestrictChatMember(userID, chatID, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// PromoteChatMember performs promoteChatMember telegram api method.
//
// options:
//
// option.CanChangeInfo(), option.CanPostMessages(), option.CanDeleteMessages(), option.CanEditMessages()
// option.CanInviteUsers(), option.CanRestrictUsers(), option.CanPinMessages(), option.CanPromoteMembers()
//
// Reference: https://core.telegram.org/bots/api#promotechatmember
// TODO: rework true/false
func (t *Client) PromoteChatMember(ctx context.Context, userID int, chatID int64, opts ...method.Option[*method.PromoteChatMember]) (bool, error) {
	m := method.NewPromoteChatMember(userID, chatID, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
