package telegram_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"maps"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/irisked/gogram/telegram"
)

// request is what the fake server saw.
type request struct {
	calls  int
	method string
	json   map[string]any    // body of a JSON request
	fields map[string]string // form fields of a multipart request
	files  map[string]string // part name -> "filename:content"
}

// newServer starts a fake Bot API server that records each request and
// answers with result.
func newServer(t *testing.T, result string) (*telegram.Client, *request) {
	t.Helper()
	got := &request{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.calls++
		got.method = r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		mediaType, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		switch mediaType {
		case "application/json":
			if err := json.NewDecoder(r.Body).Decode(&got.json); err != nil {
				t.Errorf("decode JSON body: %v", err)
			}
		case "multipart/form-data":
			got.fields, got.files = map[string]string{}, map[string]string{}
			mr := multipart.NewReader(r.Body, params["boundary"])
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Errorf("read multipart: %v", err)
					break
				}
				b, _ := io.ReadAll(p)
				if p.FileName() != "" {
					got.files[p.FormName()] = p.FileName() + ":" + string(b)
				} else {
					got.fields[p.FormName()] = string(b)
				}
			}
		default:
			t.Errorf("unexpected Content-Type %q", mediaType)
		}
		io.WriteString(w, `{"ok":true,"result":`+result+`}`)
	}))
	t.Cleanup(srv.Close)
	return telegram.New("token", telegram.WithBaseURL(srv.URL), telegram.WithHTTPClient(srv.Client())), got
}

const messageJSON = `{"message_id":7,"date":0,"chat":{"id":42,"type":"private"}}`

var message = &telegram.Message{MessageID: 7, Chat: &telegram.Chat{ID: 42, Type: "private"}}

func file(name, content string) *telegram.InputFile {
	return telegram.FileReader(name, strings.NewReader(content))
}

func TestRequestEncoding(t *testing.T) {
	photoPath := filepath.Join(t.TempDir(), "cat.jpg")
	if err := os.WriteFile(photoPath, []byte("JPEG from disk"), 0o600); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		result     string
		call       func(ctx context.Context, c *telegram.Client) error
		wantMethod string
		// A request without uploads is JSON (wantJSON); one with uploads
		// is multipart (wantFields and wantFiles).
		wantJSON   map[string]any
		wantFields map[string]string
		wantFiles  map[string]string
	}{
		{
			name:   "options are encoded as JSON",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				kb := telegram.NewInlineKeyboard(telegram.CallbackButton("Yes", "vote", "yes"))
				_, err := c.SendMessage(ctx, 42, telegram.Markdown("*hi*"),
					telegram.WithReplyMarkup(kb),
					telegram.WithReplyToMessageID(5),
					telegram.WithDisableNotification(),
					telegram.WithDisableWebPagePreview(),
				)
				return err
			},
			wantMethod: "sendMessage",
			wantJSON: map[string]any{
				"chat_id":              float64(42),
				"text":                 "*hi*",
				"parse_mode":           "Markdown",
				"disable_notification": true,
				"link_preview_options": map[string]any{"is_disabled": true},
				"reply_parameters":     map[string]any{"message_id": float64(5)},
				"reply_markup": map[string]any{"inline_keyboard": []any{[]any{
					map[string]any{"text": "Yes", "callback_data": "vote:yes"},
				}}},
			},
		},
		{
			name:   "nil options are ignored",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendMessage(ctx, 42, telegram.Text("hi"), nil)
				return err
			},
			wantMethod: "sendMessage",
			wantJSON:   map[string]any{"chat_id": float64(42), "text": "hi"},
		},
		{
			name:   "chat addressed by username",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				req := telegram.NewSendMessage(0, telegram.Text("hi"))
				req.ChatID = telegram.ChatUsername("@mychannel")
				return c.Do(ctx, req, nil)
			},
			wantMethod: "sendMessage",
			wantJSON:   map[string]any{"chat_id": "@mychannel", "text": "hi"},
		},
		{
			name:   "file ID is sent as JSON",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendPhoto(ctx, 1, telegram.FileID("abc"),
					telegram.WithCaption(telegram.HTML("<b>x</b>")))
				return err
			},
			wantMethod: "sendPhoto",
			wantJSON:   map[string]any{"chat_id": float64(1), "photo": "abc", "caption": "<b>x</b>", "parse_mode": "HTML"},
		},
		{
			name:   "file URL is sent as JSON",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendPhoto(ctx, 1, telegram.FileURL("https://example.com/cat.jpg"))
				return err
			},
			wantMethod: "sendPhoto",
			wantJSON:   map[string]any{"chat_id": float64(1), "photo": "https://example.com/cat.jpg"},
		},
		{
			name:   "upload is sent as multipart",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				kb := telegram.NewInlineKeyboard(telegram.URLButton("Go", "https://example.com"))
				_, err := c.SendPhoto(ctx, 1, file("cat.jpg", "JPEG"),
					telegram.WithReplyMarkup(kb),
					telegram.WithHasSpoiler(),
				)
				return err
			},
			wantMethod: "sendPhoto",
			wantFields: map[string]string{
				"chat_id":      "1",
				"has_spoiler":  "true",
				"reply_markup": `{"inline_keyboard":[[{"text":"Go","url":"https://example.com"}]]}`,
			},
			wantFiles: map[string]string{"photo": "cat.jpg:JPEG"},
		},
		{
			name:   "upload from a path",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendPhoto(ctx, 1, telegram.LocalPath(photoPath))
				return err
			},
			wantMethod: "sendPhoto",
			wantFields: map[string]string{"chat_id": "1"},
			wantFiles:  map[string]string{"photo": "cat.jpg:JPEG from disk"},
		},
		{
			name:   "several uploads are sent as file parts",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendDocument(ctx, 1, file("a.txt", "hello"),
					telegram.WithThumbnail(file("t.jpg", "thumb")),
					telegram.WithCaption(telegram.Text("doc")),
				)
				return err
			},
			wantMethod: "sendDocument",
			wantFields: map[string]string{"chat_id": "1", "caption": "doc"},
			wantFiles:  map[string]string{"document": "a.txt:hello", "thumbnail": "t.jpg:thumb"},
		},
		{
			name:   "uploads inside media use attach",
			result: "[" + messageJSON + "]",
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendMediaGroup(ctx, 1, []telegram.InputMedia{
					&telegram.InputMediaPhoto{Media: file("a.jpg", "A"), Caption: "a"},
					&telegram.InputMediaPhoto{Media: telegram.FileID("existing")},
				})
				return err
			},
			wantMethod: "sendMediaGroup",
			wantFields: map[string]string{
				"chat_id": "1",
				"media":   `[{"type":"photo","media":"attach://file0","caption":"a"},{"type":"photo","media":"existing"}]`,
			},
			wantFiles: map[string]string{"file0": "a.jpg:A"},
		},
		{
			name:   "interface values carry their kind",
			result: "true",
			call: func(ctx context.Context, c *telegram.Client) error {
				return c.SetMessageReaction(ctx, 42, 7,
					telegram.WithReaction(&telegram.ReactionTypeEmoji{Emoji: "👍"}))
			},
			wantMethod: "setMessageReaction",
			wantJSON: map[string]any{
				"chat_id": float64(42), "message_id": float64(7),
				"reaction": []any{map[string]any{"type": "emoji", "emoji": "👍"}},
			},
		},
		{
			name:   "edit of a chat message",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.EditMessageText(ctx, telegram.ChatMessage(42, 7), telegram.Text("new"))
				return err
			},
			wantMethod: "editMessageText",
			wantJSON:   map[string]any{"chat_id": float64(42), "message_id": float64(7), "text": "new"},
		},
		{
			name:   "edit of an inline message",
			result: "true",
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.EditMessageText(ctx, telegram.InlineMessage("inl"), telegram.Text("new"),
					telegram.WithInlineKeyboard(telegram.NewInlineKeyboard(telegram.CallbackButton("A", "a"))))
				return err
			},
			wantMethod: "editMessageText",
			wantJSON: map[string]any{
				"inline_message_id": "inl", "text": "new",
				"reply_markup": map[string]any{"inline_keyboard": []any{[]any{
					map[string]any{"text": "A", "callback_data": "a"},
				}}},
			},
		},
		{
			name:   "Do sends any request",
			result: `{"id":1,"is_bot":true,"first_name":"bot"}`,
			call: func(ctx context.Context, c *telegram.Client) error {
				return c.Do(ctx, telegram.NewGetMe(), nil)
			},
			wantMethod: "getMe",
			wantJSON:   map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, got := newServer(t, tt.result)
			if err := tt.call(context.Background(), c); err != nil {
				t.Fatalf("call error = %v", err)
			}

			if got.method != tt.wantMethod {
				t.Errorf("method = %s, want %s", got.method, tt.wantMethod)
			}
			if !reflect.DeepEqual(got.json, tt.wantJSON) {
				g, _ := json.Marshal(got.json)
				w, _ := json.Marshal(tt.wantJSON)
				t.Errorf("JSON body\n got: %s\nwant: %s", g, w)
			}
			if !maps.Equal(got.fields, tt.wantFields) {
				t.Errorf("form fields\n got: %v\nwant: %v", got.fields, tt.wantFields)
			}
			if !maps.Equal(got.files, tt.wantFiles) {
				t.Errorf("files\n got: %v\nwant: %v", got.files, tt.wantFiles)
			}
		})
	}
}

func TestInvalidRequestsAreNotSent(t *testing.T) {
	tests := []struct {
		name    string
		call    func(ctx context.Context, c *telegram.Client) error
		wantErr string
	}{
		{
			name: "callback data too long",
			call: func(ctx context.Context, c *telegram.Client) error {
				kb := telegram.NewInlineKeyboard(telegram.CallbackButton("Big", "page", strings.Repeat("x", 60)))
				_, err := c.SendMessage(ctx, 1, telegram.Text("hi"), telegram.WithReplyMarkup(kb))
				return err
			},
			wantErr: "callback data is 65 bytes",
		},
		{
			name: "callback data too long in an edit",
			call: func(ctx context.Context, c *telegram.Client) error {
				kb := telegram.NewInlineKeyboard().Row(telegram.CallbackButton("Big", strings.Repeat("y", 65)))
				_, err := c.EditMessageReplyMarkup(ctx, telegram.ChatMessage(1, 2), telegram.WithInlineKeyboard(kb))
				return err
			},
			wantErr: "callback data is 65 bytes",
		},
		{
			name: "missing file",
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendDocument(ctx, 1, telegram.LocalPath(filepath.Join(t.TempDir(), "missing.txt")))
				return err
			},
			wantErr: "missing.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, got := newServer(t, messageJSON)
			err := tt.call(context.Background(), c)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %v, want one containing %q", err, tt.wantErr)
			}
			if got.calls != 0 {
				t.Errorf("request was sent %d times, want 0", got.calls)
			}
		})
	}
}

func TestResultDecoding(t *testing.T) {
	tests := []struct {
		name   string
		result string
		call   func(ctx context.Context, c *telegram.Client) (any, error)
		want   any
	}{
		{
			name:   "message",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.SendMessage(ctx, 42, telegram.Text("hi"))
			},
			want: message,
		},
		{
			name:   "list of messages",
			result: "[" + messageJSON + "," + messageJSON + "]",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.SendMediaGroup(ctx, 42, nil)
			},
			want: []telegram.Message{*message, *message},
		},
		{
			name:   "edited chat message",
			result: messageJSON,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.EditMessageText(ctx, telegram.ChatMessage(42, 7), telegram.Text("new"))
			},
			want: message,
		},
		{
			name:   "edited inline message is nil",
			result: "true",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.EditMessageText(ctx, telegram.InlineMessage("inl"), telegram.Text("new"))
			},
			want: (*telegram.Message)(nil),
		},
		{
			name:   "interface result",
			result: `{"status":"administrator","user":{"id":5,"is_bot":false,"first_name":"Ann"},"can_be_edited":false,"can_delete_messages":true}`,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.GetChatMember(ctx, 42, 5)
			},
			want: &telegram.ChatMemberAdministrator{User: &telegram.User{ID: 5, FirstName: "Ann"}, CanDeleteMessages: true},
		},
		{
			name:   "interface result of an unknown kind",
			result: `{"status":"wizard","user":{"id":5}}`,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.GetChatMember(ctx, 42, 5)
			},
			want: &telegram.UnknownVariant{Raw: json.RawMessage(`{"status":"wizard","user":{"id":5}}`)},
		},
		{
			name:   "list of interface results",
			result: `[{"status":"creator","user":{"id":1},"is_anonymous":false},{"status":"administrator","user":{"id":2},"can_be_edited":true}]`,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.GetChatAdministrators(ctx, 42)
			},
			want: []telegram.ChatMember{
				&telegram.ChatMemberOwner{User: &telegram.User{ID: 1}},
				&telegram.ChatMemberAdministrator{User: &telegram.User{ID: 2}, CanBeEdited: true},
			},
		},
		{
			name:   "Do decodes into the given value",
			result: `{"id":1,"is_bot":true,"first_name":"bot"}`,
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				var u telegram.User
				err := c.Do(ctx, telegram.NewGetMe(), &u)
				return u, err
			},
			want: telegram.User{ID: 1, IsBot: true, FirstName: "bot"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := newServer(t, tt.result)
			got, err := tt.call(context.Background(), c)
			if err != nil {
				t.Fatalf("call error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("result = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAPIErrors(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     telegram.Error
		wantMsg  string
	}{
		{
			name:     "rate limited",
			response: `{"ok":false,"error_code":429,"description":"Too Many Requests: retry after 3","parameters":{"retry_after":3}}`,
			want:     telegram.Error{Method: "getMe", Code: 429, Description: "Too Many Requests: retry after 3", RetryAfter: 3 * time.Second},
			wantMsg:  "telegram: getMe: 429 Too Many Requests: retry after 3",
		},
		{
			name:     "group migrated",
			response: `{"ok":false,"error_code":400,"description":"Bad Request: group chat was upgraded to a supergroup chat","parameters":{"migrate_to_chat_id":-1001}}`,
			want:     telegram.Error{Method: "getMe", Code: 400, Description: "Bad Request: group chat was upgraded to a supergroup chat", MigrateToChatID: -1001},
			wantMsg:  "telegram: getMe: 400 Bad Request: group chat was upgraded to a supergroup chat",
		},
		{
			name:     "without parameters",
			response: `{"ok":false,"error_code":403,"description":"Forbidden: bot was blocked by the user"}`,
			want:     telegram.Error{Method: "getMe", Code: 403, Description: "Forbidden: bot was blocked by the user"},
			wantMsg:  "telegram: getMe: 403 Forbidden: bot was blocked by the user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, tt.response)
			}))
			defer srv.Close()
			c := telegram.New("token", telegram.WithBaseURL(srv.URL))

			_, err := c.GetMe(context.Background())
			var apiErr *telegram.Error
			if !errors.As(err, &apiErr) {
				t.Fatalf("error = %v, want *telegram.Error", err)
			}
			if *apiErr != tt.want {
				t.Errorf("error = %+v, want %+v", *apiErr, tt.want)
			}
			if err.Error() != tt.wantMsg {
				t.Errorf("Error() = %q, want %q", err.Error(), tt.wantMsg)
			}
		})
	}
}
