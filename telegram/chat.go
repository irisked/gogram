package telegram

import (
	"context"
	"os"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SendChatAction performs sendChatAction telegram api method.
//
// Reference: https://core.telegram.org/bots/api#sendchataction
func (t *Client) SendChatAction(ctx context.Context, chatID int64, action types.ChatAction) (types.Message, error) {
	m := method.NewSendChatAction(chatID, action)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// ExportChatInviteLink performs exportChatInviteLink telegram api method.
//
// Reference: https://core.telegram.org/bots/api#exportchatinvitelink
func (t *Client) ExportChatInviteLink(ctx context.Context, chatID int64) (string, error) {
	m := method.NewExportChatInviteLink(chatID)
	var result string
	err := t.process(ctx, m, &result)
	return result, err
}

// SetChatPhoto performs setChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatphoto
func (t *Client) SetChatPhoto(ctx context.Context, chatID int64, file *os.File) (bool, error) {
	m := method.NewSetChatPhoto(chatID, file)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteChatPhoto performs deleteChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatphoto
func (t *Client) DeleteChatPhoto(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewDeleteChatPhoto(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetChatTitle performs setChatTitle telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchattitle
func (t *Client) SetChatTitle(ctx context.Context, chatID int64, title string) (bool, error) {
	m := method.NewSetChatTitle(chatID, title)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetChatDescription performs setChatDescription telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatdescription
func (t *Client) SetChatDescription(ctx context.Context, chatID int64, description string) (bool, error) {
	m := method.NewSetChatDescription(chatID, description)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// PinChatMessage performs pinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#pinchatmessage
func (t *Client) PinChatMessage(ctx context.Context, chatID int64, messageID int, opts ...method.Option[*method.PinChatMessage]) (bool, error) {
	m := method.NewPinChatMessage(chatID, messageID, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// UnpinChatMessage performs unpinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinchatmessage
func (t *Client) UnpinChatMessage(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewUnpinChatMessage(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// LeaveChat performs leaveChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#leavechat
func (t *Client) LeaveChat(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewLeaveChat(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetChat performs getChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchat
func (t *Client) GetChat(ctx context.Context, chatID int64) (types.Chat, error) {
	m := method.NewGetChat(chatID)
	chat := types.Chat{}
	err := t.process(ctx, m, &chat)
	return chat, err
}

// GetChatAdministrators performs getChatAdministrators telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatadministrators
// TODO: results is array
func (t *Client) GetChatAdministrators(ctx context.Context, chatID int64) (types.ChatMember, error) {
	m := method.NewGetChatAdministrators(chatID)
	administrators := types.ChatMember{}
	err := t.process(ctx, m, &administrators)
	return administrators, err
}

// GetChatMembersCount performs getChatMembersCount telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmemberscount
func (t *Client) GetChatMembersCount(ctx context.Context, chatID int64) (int, error) {
	m := method.NewGetChatMembersCount(chatID)
	var result int
	err := t.process(ctx, m, &result)
	return result, err
}

// GetChatMember performs getChatMember telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmember
func (t *Client) GetChatMember(ctx context.Context, chatID int64, userID int) (types.ChatMember, error) {
	m := method.NewGetChatMember(chatID, userID)
	member := types.ChatMember{}
	err := t.process(ctx, m, &member)
	return member, err
}

// SetChatStickerSet performs setChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatstickerset
func (t *Client) SetChatStickerSet(ctx context.Context, chatID int64, stickerSetName string) (bool, error) {
	m := method.NewSetChatStickerSet(chatID, stickerSetName)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteChatStickerSet performs deleteChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatstickerset
func (t *Client) DeleteChatStickerSet(ctx context.Context, chatID int64) (bool, error) {
	m := method.NewDeleteChatStickerSet(chatID)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
