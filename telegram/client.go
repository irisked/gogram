package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/irisked/gogram/telegram/internal/transport"
)

// Request is implemented by every Bot API request type in this package
// (SendMessage, GetUpdates, ...). Method returns the API method name,
// e.g. "sendMessage".
//
// Implement Request yourself to call a Bot API method this package does
// not cover yet, and pass it to [Client.Do].
type Request interface {
	Method() string
}

// Client is a Telegram Bot API client. It is safe for concurrent use.
type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
	transport  *transport.Client
}

// ClientOption configures a Client created by New.
type ClientOption func(*Client)

// WithHTTPClient sets the HTTP client used for API calls. Its Timeout,
// if any, must be longer than the long-polling timeout used with
// GetUpdates; prefer per-call deadlines through context instead.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) { c.httpClient = hc }
}

// WithBaseURL points the client at a different Bot API server, such as
// a self-hosted telegram-bot-api instance or a test server. The default
// is https://api.telegram.org.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) { c.baseURL = url }
}

// New creates a Client for the given bot token.
func New(token string, opts ...ClientOption) *Client {
	c := &Client{token: token}
	for _, o := range opts {
		o(c)
	}
	c.transport = transport.New(c.token, c.baseURL, c.httpClient)
	return c
}

// Do sends any Request and decodes its result into result, which must
// be a pointer (or nil to discard the result). Failed calls return an
// *Error.
func (c *Client) Do(ctx context.Context, req Request, result any) error {
	body, contentType, err := encodeRequest(req)
	if err != nil {
		return fmt.Errorf("telegram: encode %s: %w", req.Method(), err)
	}
	res, err := c.transport.Call(ctx, req.Method(), contentType, body)
	if err != nil {
		return fmt.Errorf("telegram: %s: %w", req.Method(), err)
	}
	if !res.OK {
		return newError(req.Method(), res)
	}
	if result == nil {
		return nil
	}
	if err := json.Unmarshal(res.Result, result); err != nil {
		return fmt.Errorf("telegram: decode %s result: %w", req.Method(), err)
	}
	return nil
}

// FileURL returns the URL to download a file. filePath is File.FilePath
// as returned by GetFile. The URL embeds the bot token: don't log or
// share it.
func (c *Client) FileURL(filePath string) string {
	return c.transport.FileURL(filePath)
}

// processEdit handles methods that return the edited Message for chat
// messages but just true for inline messages.
func (c *Client) processEdit(ctx context.Context, req Request) (*Message, error) {
	var raw json.RawMessage
	if err := c.Do(ctx, req, &raw); err != nil {
		return nil, err
	}
	if bytes.Equal(bytes.TrimSpace(raw), []byte("true")) {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode %s result: %w", req.Method(), err)
	}
	return &m, nil
}
