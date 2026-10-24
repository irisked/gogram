package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// GetMe performs getMe telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getme
func (t *Client) GetMe(ctx context.Context) (types.User, error) {
	m := method.NewGetMe()
	user := types.User{}
	err := t.process(ctx, m, &user)
	return user, err
}

// GetUpdates performs getUpdates telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getupdates
func (t *Client) GetUpdates(ctx context.Context, opts ...method.Option[*method.GetUpdates]) ([]types.Update, error) {
	m := method.NewGetUpdates(opts...)
	var updates []types.Update
	err := t.process(ctx, m, &updates)
	return updates, err
}
