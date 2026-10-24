package sandbox_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/irisked/gogram/sandbox"
	"github.com/irisked/gogram/telegram"
)

func TestServerAnswers(t *testing.T) {
	tests := []struct {
		name  string
		setup func(s *sandbox.Server)
		call  func(ctx context.Context, c *telegram.Client) (any, error)
		check func(t *testing.T, result any, err error)
	}{
		{
			name: "getMe returns the test bot",
			call: func(ctx context.Context, c *telegram.Client) (any, error) { return c.GetMe(ctx) },
			check: func(t *testing.T, result any, err error) {
				if me, _ := result.(*telegram.User); err != nil || me.Username != sandbox.BotUsername {
					t.Errorf("GetMe() = %v, %v", result, err)
				}
			},
		},
		{
			name: "send methods return a message in the chat",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.SendMessage(ctx, sandbox.GroupID, telegram.Text("hi"))
			},
			check: func(t *testing.T, result any, err error) {
				m, _ := result.(*telegram.Message)
				if err != nil || m.Chat.ID != sandbox.GroupID || m.Text != "hi" || m.MessageID == 0 || !m.IsAccessible() {
					t.Errorf("SendMessage() = %+v, %v", m, err)
				}
			},
		},
		{
			name: "media groups return one message per item",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.SendMediaGroup(ctx, 1, []telegram.InputMedia{
					&telegram.InputMediaPhoto{Media: telegram.FileID("a")},
					&telegram.InputMediaPhoto{Media: telegram.FileReader("b.jpg", strings.NewReader("B"))},
				})
			},
			check: func(t *testing.T, result any, err error) {
				if msgs, _ := result.([]telegram.Message); err != nil || len(msgs) != 2 {
					t.Errorf("SendMediaGroup() = %v, %v", result, err)
				}
			},
		},
		{
			name: "edits of inline messages return nil",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.EditMessageText(ctx, telegram.InlineMessage("x"), telegram.Text("new"))
			},
			check: func(t *testing.T, result any, err error) {
				if m, _ := result.(*telegram.Message); err != nil || m != nil {
					t.Errorf("EditMessageText() = %v, %v", result, err)
				}
			},
		},
		{
			name: "edits of chat messages return the message",
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.EditMessageText(ctx, telegram.ChatMessage(5, 9), telegram.Text("new"))
			},
			check: func(t *testing.T, result any, err error) {
				if m, _ := result.(*telegram.Message); err != nil || m.MessageID != 9 || m.Chat.ID != 5 {
					t.Errorf("EditMessageText() = %+v, %v", m, err)
				}
			},
		},
		{
			name: "custom answer",
			setup: func(s *sandbox.Server) {
				s.Handle("getChatMemberCount", func(sandbox.Request) (any, error) { return 42, nil })
			},
			call: func(ctx context.Context, c *telegram.Client) (any, error) { return c.GetChatMemberCount(ctx, 1) },
			check: func(t *testing.T, result any, err error) {
				if n, _ := result.(int); err != nil || n != 42 {
					t.Errorf("GetChatMemberCount() = %v, %v", result, err)
				}
			},
		},
		{
			name: "custom API error",
			setup: func(s *sandbox.Server) {
				s.Handle("sendMessage", func(sandbox.Request) (any, error) {
					return nil, &telegram.Error{Code: 429, Description: "Too Many Requests", RetryAfter: 3 * time.Second}
				})
			},
			call: func(ctx context.Context, c *telegram.Client) (any, error) {
				return c.SendMessage(ctx, 1, telegram.Text("hi"))
			},
			check: func(t *testing.T, _ any, err error) {
				apiErr, ok := errors.AsType[*telegram.Error](err)
				if !ok || apiErr.Code != 429 || apiErr.RetryAfter != 3*time.Second {
					t.Errorf("error = %v, want a 429 with retry after 3s", err)
				}
			},
		},
		{
			name: "other errors",
			setup: func(s *sandbox.Server) {
				s.Handle("getMe", func(sandbox.Request) (any, error) { return nil, errors.New("down") })
			},
			call: func(ctx context.Context, c *telegram.Client) (any, error) { return c.GetMe(ctx) },
			check: func(t *testing.T, _ any, err error) {
				if apiErr, ok := errors.AsType[*telegram.Error](err); !ok || apiErr.Code != 500 || apiErr.Description != "down" {
					t.Errorf("error = %v, want a 500 saying down", err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sandbox.NewServer(t)
			if tt.setup != nil {
				tt.setup(s)
			}
			result, err := tt.call(t.Context(), s.Client())
			tt.check(t, result, err)
		})
	}
}

func TestServerRecordsRequests(t *testing.T) {
	tests := []struct {
		name       string
		call       func(ctx context.Context, c *telegram.Client) error
		wantMethod string
		wantString map[string]string
		wantInt    map[string]int64
		wantFiles  map[string]string // part name -> "filename:content"
	}{
		{
			name: "JSON request",
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendMessage(ctx, 9007199254740993, telegram.HTML("<b>hi</b>"),
					telegram.WithReplyToMessageID(4))
				return err
			},
			wantMethod: "sendMessage",
			wantString: map[string]string{"text": "<b>hi</b>", "parse_mode": "HTML", "reply_parameters": `{"message_id":4}`, "missing": ""},
			wantInt:    map[string]int64{"chat_id": 9007199254740993, "missing": 0},
		},
		{
			name: "multipart request",
			call: func(ctx context.Context, c *telegram.Client) error {
				_, err := c.SendDocument(ctx, 7, telegram.FileReader("a.txt", strings.NewReader("hello")),
					telegram.WithCaption(telegram.Text("doc")), telegram.WithDisableNotification())
				return err
			},
			wantMethod: "sendDocument",
			wantString: map[string]string{"caption": "doc", "disable_notification": "true"},
			wantInt:    map[string]int64{"chat_id": 7},
			wantFiles:  map[string]string{"document": "a.txt:hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sandbox.NewServer(t)
			if err := tt.call(t.Context(), s.Client()); err != nil {
				t.Fatal(err)
			}
			calls := s.Calls(tt.wantMethod)
			if len(calls) != 1 {
				t.Fatalf("calls = %v, want one %s", s.Requests(), tt.wantMethod)
			}
			r := calls[0]
			for k, want := range tt.wantString {
				if got := r.String(k); got != want {
					t.Errorf("String(%q) = %q, want %q", k, got, want)
				}
			}
			for k, want := range tt.wantInt {
				if got := r.Int(k); got != want {
					t.Errorf("Int(%q) = %d, want %d", k, got, want)
				}
			}
			for k, want := range tt.wantFiles {
				if f := r.Files[k]; f.Name+":"+string(f.Content) != want {
					t.Errorf("file %s = %+v, want %s", k, f, want)
				}
			}
		})
	}
}

func TestServerGetUpdates(t *testing.T) {
	tests := []struct {
		name    string
		updates []telegram.Update
		offset  int
		limit   int
		wantIDs []int
	}{
		{"nothing queued", nil, 0, 0, nil},
		{"IDs are assigned in order", []telegram.Update{sandbox.Message("a"), sandbox.Message("b")}, 0, 0, []int{1, 2}},
		{"offset skips handled updates", []telegram.Update{sandbox.Message("a"), sandbox.Message("b")}, 2, 0, []int{2}},
		{"limit", []telegram.Update{sandbox.Message("a"), sandbox.Message("b"), sandbox.Message("c")}, 0, 2, []int{1, 2}},
		{"explicit IDs are kept", []telegram.Update{{UpdateID: 10, Message: sandbox.Message("a").Message}, sandbox.Message("b")}, 0, 0, []int{10, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sandbox.NewServer(t)
			s.AddUpdates(tt.updates...)
			opts := []telegram.GetUpdatesOption{telegram.WithOffset(tt.offset)}
			if tt.limit > 0 {
				opts = append(opts, telegram.WithLimit(tt.limit))
			}
			got, err := s.Client().GetUpdates(t.Context(), opts...)
			if err != nil {
				t.Fatal(err)
			}
			var ids []int
			for _, u := range got {
				ids = append(ids, u.UpdateID)
			}
			if !equal(ids, tt.wantIDs) {
				t.Errorf("update IDs = %v, want %v", ids, tt.wantIDs)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestUpdateBuilders(t *testing.T) {
	tests := []struct {
		name        string
		update      telegram.Update
		wantType    telegram.UpdateType
		wantCommand string
		wantChat    int64
	}{
		{"message", sandbox.Message("hello"), telegram.UpdateMessage, "", sandbox.ChatID},
		{"message that is a command", sandbox.Message("/help me"), telegram.UpdateMessage, "help", sandbox.ChatID},
		{"command without slash", sandbox.Command("start"), telegram.UpdateMessage, "start", sandbox.ChatID},
		{"group message", sandbox.GroupMessage("/ban@test_bot 1"), telegram.UpdateMessage, "ban", sandbox.GroupID},
		{"callback", sandbox.Callback("vote:up"), telegram.UpdateCallbackQuery, "", sandbox.ChatID},
		{"inline query", sandbox.InlineQuery("cats"), telegram.UpdateInlineQuery, "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.update.Type(); got != tt.wantType {
				t.Errorf("Type() = %q, want %q", got, tt.wantType)
			}
			var command string
			var chat int64
			switch {
			case tt.update.Message != nil:
				command, chat = tt.update.Message.Command(), tt.update.Message.Chat.ID
			case tt.update.CallbackQuery != nil:
				chat = tt.update.CallbackQuery.Message.Chat.ID
			}
			if command != tt.wantCommand || chat != tt.wantChat {
				t.Errorf("command, chat = %q, %d; want %q, %d", command, chat, tt.wantCommand, tt.wantChat)
			}
		})
	}
}
