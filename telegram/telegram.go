// Package telegram implements a client for the Telegram Bot API.
//
// Client's methods are grouped by category across several files
// mirroring the categorization already used by the method package
// (chat.go, media.go, sticker.go, forum-topic.go, and so on).
package telegram

import (
	"context"
	"encoding/json"

	"github.com/irisked/gogram/telegram/internal/transport"
	"github.com/irisked/gogram/telegram/method"
)

// Client is a Telegram Bot API client. The underlying *transport.Client
// (telegram/internal/transport) is an unrelated, internal HTTP client --
// not to be confused with this exported Client.
type Client struct {
	client *transport.Client
}

// New creates a Client for the given bot token.
func New(token string) *Client {
	c := new(Client)
	c.client = transport.NewClient(token)
	return c
}

func (t *Client) process(ctx context.Context, m method.TelegramMethod, result interface{}) error {
	response, err := t.client.Call(ctx, m)
	if err != nil {
		return err
	}
	return json.Unmarshal(response.Result, result)
}

// TODO: telegram passport methods.
// TODO: games methods.
