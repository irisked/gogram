package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SendLocation performs sendLocation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendlocation
func (t *Client) SendLocation(ctx context.Context, chatID int64, location types.Location, opts ...method.Option[*method.SendLocation]) (types.Message, error) {
	m := method.NewSendLocation(chatID, location, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendVenue performs sendVenue telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard), option.Foursquare(id, type string)
//
// Reference: https://core.telegram.org/bots/api#sendvenue
func (t *Client) SendVenue(ctx context.Context, chatID int64, venue types.Venue, opts ...method.Option[*method.SendVenue]) (types.Message, error) {
	m := method.NewSendVenue(chatID, venue, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendContact performs sendContact telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendcontact
func (t *Client) SendContact(ctx context.Context, chatID int64, contact types.Contact, opts ...method.Option[*method.SendContact]) (types.Message, error) {
	m := method.NewSendContact(chatID, contact, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}
