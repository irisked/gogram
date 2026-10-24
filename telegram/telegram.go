// Package telegram implements a client for the Telegram Bot API.
//
// Client's methods are grouped by category across several files
// mirroring the categorization already used by the method package
// (chat.go, media.go, sticker.go, forum-topic.go, and so on).
package telegram

import (
	"context"
	"encoding/json"

	"github.com/irisked/gogram/telegram/internal/net"
	"github.com/irisked/gogram/telegram/method"
)

// Client is a Telegram Bot API client. The underlying *net.Client
// (telegram/internal/net) is an unrelated, internal HTTP transport --
// not to be confused with this exported Client.
type Client struct {
	client *net.Client
}

// New creates a Client for the given bot token.
func New(token string) *Client {
	c := new(Client)
	c.client = net.NewClient(token)
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
