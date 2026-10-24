package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SetMyCommands performs setMyCommands telegram api method.
//
// options:
//
// option.SetMyCommandsScope(), option.SetMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmycommands
func (t *Client) SetMyCommands(ctx context.Context, commands []method.BotCommand, opts ...method.Option[*method.SetMyCommands]) (bool, error) {
	m := method.NewSetMyCommands(commands, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetMyCommands performs getMyCommands telegram api method.
//
// options:
//
// option.GetMyCommandsScope(), option.GetMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmycommands
func (t *Client) GetMyCommands(ctx context.Context, opts ...method.Option[*method.GetMyCommands]) ([]method.BotCommand, error) {
	m := method.NewGetMyCommands(opts...)
	var result []method.BotCommand
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteMyCommands performs deleteMyCommands telegram api method.
//
// options:
//
// option.DeleteMyCommandsScope(), option.DeleteMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#deletemycommands
func (t *Client) DeleteMyCommands(ctx context.Context, opts ...method.Option[*method.DeleteMyCommands]) (bool, error) {
	m := method.NewDeleteMyCommands(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetMyName performs setMyName telegram api method.
//
// options:
//
// option.SetMyNameName(), option.SetMyNameLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmyname
func (t *Client) SetMyName(ctx context.Context, opts ...method.Option[*method.SetMyName]) (bool, error) {
	m := method.NewSetMyName(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetMyName performs getMyName telegram api method.
//
// options:
//
// option.GetMyNameLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmyname
func (t *Client) GetMyName(ctx context.Context, opts ...method.Option[*method.GetMyName]) (types.BotName, error) {
	m := method.NewGetMyName(opts...)
	var result types.BotName
	err := t.process(ctx, m, &result)
	return result, err
}

// SetMyDescription performs setMyDescription telegram api method.
//
// options:
//
// option.SetMyDescriptionDescription(), option.SetMyDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmydescription
func (t *Client) SetMyDescription(ctx context.Context, opts ...method.Option[*method.SetMyDescription]) (bool, error) {
	m := method.NewSetMyDescription(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetMyDescription performs getMyDescription telegram api method.
//
// options:
//
// option.GetMyDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmydescription
func (t *Client) GetMyDescription(ctx context.Context, opts ...method.Option[*method.GetMyDescription]) (types.BotDescription, error) {
	m := method.NewGetMyDescription(opts...)
	var result types.BotDescription
	err := t.process(ctx, m, &result)
	return result, err
}

// SetMyShortDescription performs setMyShortDescription telegram api method.
//
// options:
//
// option.SetMyShortDescriptionShortDescription(), option.SetMyShortDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmyshortdescription
func (t *Client) SetMyShortDescription(ctx context.Context, opts ...method.Option[*method.SetMyShortDescription]) (bool, error) {
	m := method.NewSetMyShortDescription(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetMyShortDescription performs getMyShortDescription telegram api method.
//
// options:
//
// option.GetMyShortDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmyshortdescription
func (t *Client) GetMyShortDescription(ctx context.Context, opts ...method.Option[*method.GetMyShortDescription]) (types.BotShortDescription, error) {
	m := method.NewGetMyShortDescription(opts...)
	var result types.BotShortDescription
	err := t.process(ctx, m, &result)
	return result, err
}

// SetMyDefaultAdministratorRights performs setMyDefaultAdministratorRights telegram api method.
//
// options:
//
// option.SetMyDefaultAdministratorRightsForSupergroups(), option.SetMyDefaultAdministratorRightsForChannels()
//
// Reference: https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (t *Client) SetMyDefaultAdministratorRights(ctx context.Context, rights method.BotAdministratorRights, opts ...method.Option[*method.SetMyDefaultAdministratorRights]) (bool, error) {
	m := method.NewSetMyDefaultAdministratorRights(&rights, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetMyDefaultAdministratorRights performs getMyDefaultAdministratorRights telegram api method.
//
// options:
//
// option.GetMyDefaultAdministratorRightsForSupergroups(), option.GetMyDefaultAdministratorRightsForChannels()
//
// Reference: https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (t *Client) GetMyDefaultAdministratorRights(ctx context.Context, opts ...method.Option[*method.GetMyDefaultAdministratorRights]) (method.BotAdministratorRights, error) {
	m := method.NewGetMyDefaultAdministratorRights(opts...)
	var result method.BotAdministratorRights
	err := t.process(ctx, m, &result)
	return result, err
}

// SetChatMenuButton performs setChatMenuButton telegram api method.
//
// options:
//
// option.SetChatMenuButtonChatID(), option.SetChatMenuButtonMenuButton()
//
// Reference: https://core.telegram.org/bots/api#setchatmenubutton
func (t *Client) SetChatMenuButton(ctx context.Context, opts ...method.Option[*method.SetChatMenuButton]) (bool, error) {
	m := method.NewSetChatMenuButton(opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetChatMenuButton performs getChatMenuButton telegram api method.
//
// options:
//
// option.GetChatMenuButtonChatID()
//
// Reference: https://core.telegram.org/bots/api#getchatmenubutton
func (t *Client) GetChatMenuButton(ctx context.Context, opts ...method.Option[*method.GetChatMenuButton]) (types.MenuButton, error) {
	m := method.NewGetChatMenuButton(opts...)
	var result types.MenuButton
	err := t.process(ctx, m, &result)
	return result, err
}
