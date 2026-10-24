package transport

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCall(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		contentType string
		body        string
		status      int
		response    string
		want        Response
		wantErr     bool
	}{
		{
			name:        "ok",
			method:      "getMe",
			contentType: "application/json",
			body:        `{}`,
			status:      http.StatusOK,
			response:    `{"ok":true,"result":{"id":1}}`,
			want:        Response{OK: true, Result: json.RawMessage(`{"id":1}`)},
		},
		{
			name:        "content type is passed through",
			method:      "sendPhoto",
			contentType: "multipart/form-data; boundary=x",
			body:        "--x--",
			status:      http.StatusOK,
			response:    `{"ok":true,"result":true}`,
			want:        Response{OK: true, Result: json.RawMessage(`true`)},
		},
		{
			name:        "rate limited",
			method:      "sendMessage",
			contentType: "application/json",
			body:        `{}`,
			status:      http.StatusTooManyRequests,
			response:    `{"ok":false,"error_code":429,"description":"Too Many Requests","parameters":{"retry_after":5}}`,
			want: Response{ErrorCode: 429, Description: "Too Many Requests",
				Parameters: &ResponseParameters{RetryAfter: 5}},
		},
		{
			name:        "group migrated",
			method:      "sendMessage",
			contentType: "application/json",
			body:        `{}`,
			status:      http.StatusBadRequest,
			response:    `{"ok":false,"error_code":400,"description":"Bad Request: group chat was upgraded","parameters":{"migrate_to_chat_id":-100}}`,
			want: Response{ErrorCode: 400, Description: "Bad Request: group chat was upgraded",
				Parameters: &ResponseParameters{MigrateToChatID: -100}},
		},
		{
			name:        "invalid response",
			method:      "getMe",
			contentType: "application/json",
			body:        `{}`,
			status:      http.StatusBadGateway,
			response:    `<html>bad gateway</html>`,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("HTTP method = %s, want POST", r.Method)
				}
				if want := "/bottest-token/" + tt.method; r.URL.Path != want {
					t.Errorf("path = %s, want %s", r.URL.Path, want)
				}
				if ct := r.Header.Get("Content-Type"); ct != tt.contentType {
					t.Errorf("Content-Type = %s, want %s", ct, tt.contentType)
				}
				if body, _ := io.ReadAll(r.Body); string(body) != tt.body {
					t.Errorf("body = %s, want %s", body, tt.body)
				}
				w.WriteHeader(tt.status)
				io.WriteString(w, tt.response)
			}))
			defer srv.Close()

			c := New("test-token", srv.URL, srv.Client())
			got, err := c.Call(context.Background(), tt.method, tt.contentType, strings.NewReader(tt.body))
			if (err != nil) != tt.wantErr {
				t.Fatalf("Call() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Call() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestCallErrorsDoNotLeakToken(t *testing.T) {
	const token = "secret-token"

	tests := []struct {
		name    string
		baseURL func(t *testing.T) string
		ctx     func() context.Context
		wantIs  error
	}{
		{
			name:    "connection refused",
			baseURL: func(*testing.T) string { return "http://127.0.0.1:1" },
		},
		{
			name: "cancelled context",
			baseURL: func(t *testing.T) string {
				srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
				t.Cleanup(srv.Close)
				return srv.URL
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			wantIs: context.Canceled,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			c := New(token, tt.baseURL(t), nil)
			_, err := c.Call(ctx, "getMe", "application/json", strings.NewReader(`{}`))
			if err == nil {
				t.Fatal("Call() error = nil, want an error")
			}
			if strings.Contains(err.Error(), token) {
				t.Errorf("error leaks the bot token: %v", err)
			}
			if tt.wantIs != nil && !errors.Is(err, tt.wantIs) {
				t.Errorf("error = %v, want it to wrap %v", err, tt.wantIs)
			}
		})
	}
}

func TestCallClosesBodyOfUnsentRequests(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
	}{
		{"invalid base URL", "http://bad host"},
		{"control character in base URL", "http://example.com/\x7f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &closeRecorder{Reader: strings.NewReader("{}")}
			_, err := New("secret-token", tt.baseURL, nil).Call(context.Background(), "getMe", "application/json", body)
			if err == nil {
				t.Fatal("Call() error = nil, want an invalid URL error")
			}
			if !body.closed {
				t.Error("body was not closed")
			}
			if strings.Contains(err.Error(), "secret-token") {
				t.Errorf("error leaks the bot token: %v", err)
			}
		})
	}
}

type closeRecorder struct {
	io.Reader
	closed bool
}

func (c *closeRecorder) Close() error {
	c.closed = true
	return nil
}

func TestFileURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		path    string
		want    string
	}{
		{"default server", "", "photos/a.jpg", "https://api.telegram.org/file/bottok/photos/a.jpg"},
		{"custom server", "http://localhost:8081", "doc.pdf", "http://localhost:8081/file/bottok/doc.pdf"},
		{"trailing slash is trimmed", "http://localhost:8081/", "doc.pdf", "http://localhost:8081/file/bottok/doc.pdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New("tok", tt.baseURL, nil)
			if got := c.FileURL(tt.path); got != tt.want {
				t.Errorf("FileURL(%q) = %s, want %s", tt.path, got, tt.want)
			}
		})
	}
}
