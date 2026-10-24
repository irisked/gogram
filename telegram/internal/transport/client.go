// Package transport sends already-encoded requests to the Telegram Bot
// API over HTTP and decodes the response envelope. Encoding request
// bodies and interpreting results is the telegram package's job.
package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DefaultBaseURL is the public Bot API server.
const DefaultBaseURL = "https://api.telegram.org"

// Response is the envelope every Bot API call returns. Result is kept
// raw so the caller can decode it into the method-specific type.
type Response struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// ResponseParameters carries extra information about a failed call.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int   `json:"retry_after"`
}

// Client posts requests to the Bot API.
type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

// New creates a Client. An empty baseURL means DefaultBaseURL and a nil
// httpClient means http.DefaultClient.
func New(token, baseURL string, httpClient *http.Client) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		token:      token,
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}

// Call posts body to the given API method and decodes the response
// envelope. A non-OK envelope is not an error at this level; the caller
// inspects Response.OK.
func (c *Client) Call(ctx context.Context, method, contentType string, body io.Reader) (Response, error) {
	url := fmt.Sprintf("%s/bot%s/%s", c.baseURL, c.token, method)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		// The client closes the body of requests it sends; this one is never
		// sent, and a streamed body would otherwise block its writer.
		if c, ok := body.(io.Closer); ok {
			c.Close()
		}
		return Response{}, redactToken(err, c.token)
	}
	req.Header.Set("Content-Type", contentType)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Response{}, redactToken(err, c.token)
	}
	defer res.Body.Close()

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("decode %s response (HTTP %d): %w", method, res.StatusCode, err)
	}
	return response, nil
}

// FileURL returns the download URL for a file path obtained from getFile.
func (c *Client) FileURL(filePath string) string {
	return fmt.Sprintf("%s/file/bot%s/%s", c.baseURL, c.token, filePath)
}

// redactToken keeps the bot token out of error messages: net/http
// includes the full request URL in *url.Error.
func redactToken(err error, token string) error {
	if token == "" || !strings.Contains(err.Error(), token) {
		return err
	}
	return &redactedError{msg: strings.ReplaceAll(err.Error(), token, "<token>"), err: err}
}

type redactedError struct {
	msg string
	err error
}

func (e *redactedError) Error() string { return e.msg }
func (e *redactedError) Unwrap() error { return e.err }
