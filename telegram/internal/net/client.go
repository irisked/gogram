package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/irisked/gogram/telegram/method"
	"io"
	"net/http"
	"time"
)

const (
	// TelegramAPIEndpoint is the endpoint for all API method with formatting for Sprintf.
	TelegramAPIEndpoint = "https://api.telegram.org/bot%s/%s"
	// TelegramFileEndpoint is the endpoint for downloading a file from Telegram.
	TelegramFileEndpoint = "https://api.telegram.org/file/bot%s/%s"
	// TelegramHTTPMethod http method for request
	TelegramHTTPMethod = "POST"
)

// Response is a response from Telegram API
// with the result stored raw
type Response struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
	RetryAfter      int   `json:"retry_after"`        // optional
}

func (rp *ResponseParameters) String() string {
	return fmt.Sprintf(string(rp.MigrateToChatID), " ", string(rp.RetryAfter))
}

// Client represents a telegram api client.
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient creates Client.
func NewClient(token string) *Client {
	client := new(Client)
	client.token = token
	client.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}

// Token returns telegram token
func (c *Client) Token() string {
	return c.token
}

// Call builds http request and performs its from passed telegram method.
func (c *Client) Call(method method.TelegramMethod) (Response, error) {
	req, err := c.buildRequest(method)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	return c.parse(res)
}

// BuildRequest returns *http.Request from method.TelegramMethod
// to be processed by client.
func (c *Client) buildRequest(method method.TelegramMethod) (*http.Request, error) {
	if method.IsMultipart() {
		return c.buildMultipartRequest(method)
	}
	return c.buildJSONRequest(method)
}

func (c *Client) buildMultipartRequest(method method.TelegramMethod) (*http.Request, error) {
	buffer, contentType, err := ToForm(method)
	if err != nil {
		return nil, err
	}
	return c.createHTTPRequest(method.Method(), buffer, contentType)
}

func (c *Client) buildJSONRequest(method method.TelegramMethod) (*http.Request, error) {
	buffer, contentType, err := ToJSON(method)
	fmt.Println(buffer)
	if err != nil {
		return nil, err
	}
	return c.createHTTPRequest(method.Method(), buffer, contentType)
}

func (c *Client) createHTTPRequest(method string, body *bytes.Buffer, content string) (*http.Request, error) {
	httpMethod := fmt.Sprintf(TelegramAPIEndpoint, c.token, method)
	req, err := http.NewRequest(TelegramHTTPMethod, httpMethod, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", content)
	return req, err
}

func (c *Client) parse(res *http.Response) (Response, error) {
	var response Response
	if err := c.decode(res.Body, &response); err != nil {
		return response, err
	}
	return c.check(response)
}

func (c *Client) decode(body io.Reader, res *Response) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(res)
	return err
}

func (c *Client) check(res Response) (Response, error) {
	if !res.Ok {
		parameters := ResponseParameters{}
		if res.Parameters != nil {
			parameters = *res.Parameters
		}
		return res, errors.New(res.Description + parameters.String())
	}
	return res, nil
}
