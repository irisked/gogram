package gogram_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/irisked/gogram"
	"github.com/irisked/gogram/sandbox"
	"github.com/irisked/gogram/telegram"
)

var quiet = gogram.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))

func noop(context.Context, *gogram.Update) error { return nil }

func TestRouting(t *testing.T) {
	// all registers a handler for most routes; each handler records its name.
	all := func(b *gogram.Bot, mark func(string) gogram.Handler) {
		b.OnAny(mark("any"))
		b.OnMessage(mark("message"))
		b.OnText(mark("text"))
		b.OnStart(mark("start"))
		b.OnCommand("/Settings", mark("settings"))
		b.OnCallbackQuery(mark("callback"))
		b.OnCallback("vote", mark("vote"))
		b.OnInlineQuery(mark("inline"))
		b.OnVideo(mark("video"))
		b.OnPollAnswer(mark("poll answer"))
		b.OnChatJoinRequest(mark("join request"))
		b.OnBusinessMessage(mark("business"))
	}
	textOnly := func(b *gogram.Bot, mark func(string) gogram.Handler) {
		b.OnText(mark("text"))
	}

	tests := []struct {
		name     string
		register func(b *gogram.Bot, mark func(string) gogram.Handler)
		username string // "" means not known
		update   telegram.Update
		want     string // "" means no handler ran
	}{
		{"text", all, "test_bot", sandbox.Message("hello"), "text"},
		{"command", all, "test_bot", sandbox.Command("/start"), "start"},
		{"command with args", all, "test_bot", sandbox.Command("/start ref42"), "start"},
		{"command is case-insensitive", all, "test_bot", sandbox.Command("/settings"), "settings"},
		{"command for this bot", all, "test_bot", sandbox.GroupMessage("/start@test_bot"), "start"},
		{"command for this bot in another case", all, "test_bot", sandbox.GroupMessage("/start@Test_Bot"), "start"},
		{"command for another bot is text", all, "test_bot", sandbox.GroupMessage("/start@other_bot"), "text"},
		{"any bot's command while the username is unknown", all, "", sandbox.GroupMessage("/start@other_bot"), "start"},
		{"command prefix is not a match", all, "test_bot", sandbox.Command("/startgame"), "message"},
		{"unregistered command falls back", all, "test_bot", sandbox.Command("/unknown"), "message"},
		{"photo falls back to message", all, "test_bot", telegram.Update{Message: &telegram.Message{Photo: []telegram.PhotoSize{{}}}}, "message"},
		{"video", all, "test_bot", telegram.Update{Message: &telegram.Message{Video: &telegram.Video{}}}, "video"},
		{"callback route", all, "test_bot", sandbox.Callback("vote:yes"), "vote"},
		{"callback route without args", all, "test_bot", sandbox.Callback("vote"), "vote"},
		{"callback route prefix is not a match", all, "test_bot", sandbox.Callback("voter:yes"), "callback"},
		{"inline query", all, "test_bot", sandbox.InlineQuery("cats"), "inline"},
		{"poll answer", all, "test_bot", telegram.Update{PollAnswer: &telegram.PollAnswer{PollID: "p"}}, "poll answer"},
		{"chat join request", all, "test_bot", telegram.Update{ChatJoinRequest: &telegram.ChatJoinRequest{}}, "join request"},
		{"business message", all, "test_bot", telegram.Update{BusinessMessage: &telegram.Message{Text: "hi"}}, "business"},
		{"edited message falls back to any", all, "test_bot", telegram.Update{EditedMessage: &telegram.Message{Text: "x"}}, "any"},
		{"unknown update type falls back to any", all, "test_bot", telegram.Update{UpdateID: 1}, "any"},
		{"unmatched update is dropped", textOnly, "test_bot", telegram.Update{UpdateID: 1}, ""},
		{"unmatched message is dropped", textOnly, "test_bot", telegram.Update{Message: &telegram.Message{Photo: []telegram.PhotoSize{{}}}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := gogram.New("token", quiet, gogram.WithUsername(tt.username))
			var got string
			tt.register(b, func(name string) gogram.Handler {
				return func(context.Context, *gogram.Update) error {
					got = name
					return nil
				}
			})

			if err := b.HandleUpdate(context.Background(), tt.update); err != nil {
				t.Fatalf("HandleUpdate() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("handled by %q, want %q", got, tt.want)
			}
		})
	}
}

type ctxKey struct{}

func TestMiddleware(t *testing.T) {
	type middleware struct {
		name  string
		block bool // don't call the next handler
	}

	tests := []struct {
		name       string
		middleware []middleware
		update     telegram.Update
		want       []string
	}{
		{
			name:   "no middleware",
			update: sandbox.Message("hi"),
			want:   []string{"handler"},
		},
		{
			name:       "single middleware",
			middleware: []middleware{{name: "log"}},
			update:     sandbox.Message("hi"),
			want:       []string{"log", "handler with value from log"},
		},
		{
			name:       "runs in the order added",
			middleware: []middleware{{name: "outer"}, {name: "inner"}},
			update:     sandbox.Message("hi"),
			want:       []string{"outer", "inner", "handler with value from inner"},
		},
		{
			name:       "can stop the chain",
			middleware: []middleware{{name: "outer"}, {name: "guard", block: true}, {name: "inner"}},
			update:     sandbox.Message("hi"),
			want:       []string{"outer", "guard"},
		},
		{
			name:       "doesn't run without a handler",
			middleware: []middleware{{name: "log"}},
			update:     sandbox.Callback("x"),
			want:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := gogram.New("token", quiet)
			var trace []string
			for _, mw := range tt.middleware {
				b.Use(func(next gogram.Handler) gogram.Handler {
					return func(ctx context.Context, u *gogram.Update) error {
						trace = append(trace, mw.name)
						if mw.block {
							return nil
						}
						return next(context.WithValue(ctx, ctxKey{}, mw.name), u)
					}
				})
			}
			b.OnText(func(ctx context.Context, u *gogram.Update) error {
				if by, ok := ctx.Value(ctxKey{}).(string); ok {
					trace = append(trace, "handler with value from "+by)
				} else {
					trace = append(trace, "handler")
				}
				return nil
			})

			if err := b.HandleUpdate(context.Background(), tt.update); err != nil {
				t.Fatalf("HandleUpdate() error = %v", err)
			}
			if !slices.Equal(trace, tt.want) {
				t.Errorf("trace = %q, want %q", trace, tt.want)
			}
		})
	}
}

var errBoom = errors.New("boom")

func TestHandleUpdateErrors(t *testing.T) {
	tests := []struct {
		name      string
		handler   gogram.Handler
		update    telegram.Update
		wantErr   error // matched with errors.Is
		wantPanic any   // the value of a *PanicError, if any
	}{
		{
			name:    "no error",
			handler: noop,
			update:  sandbox.Message("hi"),
		},
		{
			name:    "returned error",
			handler: func(context.Context, *gogram.Update) error { return fmt.Errorf("load: %w", errBoom) },
			update:  sandbox.Message("hi"),
			wantErr: errBoom,
		},
		{
			name:      "panic with a string",
			handler:   func(context.Context, *gogram.Update) error { panic("kaboom") },
			update:    sandbox.Message("hi"),
			wantPanic: "kaboom",
		},
		{
			name:      "panic with an error",
			handler:   func(context.Context, *gogram.Update) error { panic(errBoom) },
			update:    sandbox.Message("hi"),
			wantErr:   errBoom,
			wantPanic: errBoom,
		},
		{
			name:    "no handler",
			handler: func(context.Context, *gogram.Update) error { return errBoom },
			update:  sandbox.Callback("x"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := gogram.New("token", quiet)
			b.OnText(tt.handler)

			err := b.HandleUpdate(context.Background(), tt.update)
			if tt.wantErr == nil && tt.wantPanic == nil && err != nil {
				t.Fatalf("HandleUpdate() error = %v, want nil", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("HandleUpdate() error = %v, want %v", err, tt.wantErr)
			}
			if tt.wantPanic != nil {
				p, ok := errors.AsType[*gogram.PanicError](err)
				if !ok {
					t.Fatalf("HandleUpdate() error = %v, want a *PanicError", err)
				}
				if p.Value != tt.wantPanic {
					t.Errorf("panic value = %v, want %v", p.Value, tt.wantPanic)
				}
				if !bytes.Contains(p.Stack, []byte("bot_test.go")) {
					t.Errorf("stack doesn't show where the panic happened:\n%s", p.Stack)
				}
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name    string
		handler gogram.Handler
		// With a custom error handler: the error it must receive.
		wantErr error
		// With the default error handler: what the log must contain.
		wantLog []string
	}{
		{
			name:    "returned error reaches the error handler",
			handler: func(context.Context, *gogram.Update) error { return errBoom },
			wantErr: errBoom,
		},
		{
			name:    "panic reaches the error handler",
			handler: func(context.Context, *gogram.Update) error { panic(errBoom) },
			wantErr: errBoom,
		},
		{
			name:    "default handler logs the error",
			handler: func(context.Context, *gogram.Update) error { return errBoom },
			wantLog: []string{"handler failed", "boom", "update_id="},
		},
		{
			name:    "default handler logs the stack of a panic",
			handler: func(context.Context, *gogram.Update) error { panic("kaboom") },
			wantLog: []string{"handler failed", "kaboom", "stack=", "bot_test.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			var logs bytes.Buffer
			opts := []gogram.Option{gogram.WithLogger(slog.New(slog.NewTextHandler(&logs, nil)))}
			var got error
			if tt.wantErr != nil {
				opts = append(opts, gogram.WithErrorHandler(func(ctx context.Context, u *gogram.Update, err error) {
					got = err
				}))
			}
			b := srv.Bot(opts...)
			b.OnText(tt.handler)

			// Errors reach the error handler through the webhook (and polling).
			body, _ := json.Marshal(sandbox.Message("hi"))
			rec := httptest.NewRecorder()
			b.WebhookHandler("").ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body)))
			if rec.Code != http.StatusOK {
				t.Errorf("status = %d, want 200 even when the handler fails", rec.Code)
			}

			if tt.wantErr != nil && !errors.Is(got, tt.wantErr) {
				t.Errorf("error handler got %v, want %v", got, tt.wantErr)
			}
			for _, want := range tt.wantLog {
				if !strings.Contains(logs.String(), want) {
					t.Errorf("log doesn't contain %q:\n%s", want, logs.String())
				}
			}
		})
	}
}

func TestUpdateAccessors(t *testing.T) {
	pressed := sandbox.Callback("like:42,7")
	pressed.CallbackQuery.Message.Chat.ID = 9

	inline := sandbox.Callback("like:1")
	inline.CallbackQuery.Message = nil
	inline.CallbackQuery.InlineMessageID = "inl"

	tests := []struct {
		name       string
		update     telegram.Update
		wantChatID int64
		wantFromID int64 // 0 means From() is nil
		wantArgs   []string
	}{
		{"text", sandbox.Message("hello"), sandbox.ChatID, sandbox.UserID, nil},
		{"command", sandbox.Command("/ban 42  spam"), sandbox.ChatID, sandbox.UserID, []string{"42", "spam"}},
		{"command without args", sandbox.Command("/start"), sandbox.ChatID, sandbox.UserID, nil},
		{"group message", sandbox.GroupMessage("hi"), sandbox.GroupID, sandbox.UserID, nil},
		// From is the user who pressed the button, not the bot that sent the message.
		{"callback query", pressed, 9, sandbox.UserID, []string{"42", "7"}},
		{"callback query on an inline message", inline, 0, sandbox.UserID, []string{"1"}},
		{"inline query", sandbox.InlineQuery("cats"), 0, sandbox.UserID, nil},
		{"channel post", telegram.Update{ChannelPost: &telegram.Message{Text: "news", Chat: &telegram.Chat{ID: -7}}}, -7, 0, nil},
		{"chat member update", telegram.Update{MyChatMember: &telegram.ChatMemberUpdated{Chat: &telegram.Chat{ID: -8}, From: &telegram.User{ID: 6}}}, -8, 6, nil},
		{"chat join request", telegram.Update{ChatJoinRequest: &telegram.ChatJoinRequest{Chat: &telegram.Chat{ID: -9}, From: &telegram.User{ID: 7}}}, -9, 7, nil},
		{"reaction", telegram.Update{MessageReaction: &telegram.MessageReactionUpdated{Chat: &telegram.Chat{ID: -3}, User: &telegram.User{ID: 8}}}, -3, 8, nil},
		{"poll answer", telegram.Update{PollAnswer: &telegram.PollAnswer{User: &telegram.User{ID: 4}}}, 0, 4, nil},
		{"empty update", telegram.Update{UpdateID: 1}, 0, 0, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := gogram.NewUpdate(nil, tt.update)

			if got := u.ChatID(); got != tt.wantChatID {
				t.Errorf("ChatID() = %d, want %d", got, tt.wantChatID)
			}
			var fromID int64
			if from := u.From(); from != nil {
				fromID = from.ID
			}
			if fromID != tt.wantFromID {
				t.Errorf("From().ID = %d, want %d", fromID, tt.wantFromID)
			}
			if got := u.Args(); !slices.Equal(got, tt.wantArgs) {
				t.Errorf("Args() = %q, want %q", got, tt.wantArgs)
			}
		})
	}
}

func TestUpdateShortcuts(t *testing.T) {
	tests := []struct {
		name       string
		update     telegram.Update
		call       func(ctx context.Context, u *gogram.Update) error
		wantMethod string
		wantParams map[string]string
	}{
		{
			name:   "Reply",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.Reply(ctx, telegram.HTML("<b>hello</b>"), telegram.WithDisableNotification())
				return err
			},
			wantMethod: "sendMessage",
			wantParams: map[string]string{"chat_id": "100", "text": "<b>hello</b>", "parse_mode": "HTML", "disable_notification": "true"},
		},
		{
			name:   "ReplyWithPhoto",
			update: sandbox.GroupMessage("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.ReplyWithPhoto(ctx, telegram.FileID("photo-id"), telegram.WithCaption(telegram.Text("cat")))
				return err
			},
			wantMethod: "sendPhoto",
			wantParams: map[string]string{"chat_id": "-1001000000001", "photo": "photo-id", "caption": "cat"},
		},
		{
			name:   "SendChatAction",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.SendChatAction(ctx, telegram.ChatActionTyping)
			},
			wantMethod: "sendChatAction",
			wantParams: map[string]string{"chat_id": "100", "action": "typing"},
		},
		{
			name:   "BanChatMember",
			update: sandbox.GroupMessage("/ban"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.BanChatMember(ctx, 42, telegram.WithRevokeMessages())
			},
			wantMethod: "banChatMember",
			wantParams: map[string]string{"chat_id": "-1001000000001", "user_id": "42", "revoke_messages": "true"},
		},
		{
			name:   "AnswerCallbackQuery",
			update: sandbox.Callback("vote:up"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerCallbackQuery(ctx, telegram.WithText("Thanks!"))
			},
			wantMethod: "answerCallbackQuery",
			wantParams: map[string]string{"callback_query_id": "callback-1", "text": "Thanks!"},
		},
		{
			name:   "EditMessageText",
			update: sandbox.Callback("vote:up"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.EditMessageText(ctx, telegram.Text("voted"))
				return err
			},
			wantMethod: "editMessageText",
			wantParams: map[string]string{"chat_id": "100", "message_id": "1", "text": "voted"},
		},
		{
			name:   "EditMessageReplyMarkup",
			update: sandbox.Callback("vote:up"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.EditMessageReplyMarkup(ctx)
				return err
			},
			wantMethod: "editMessageReplyMarkup",
			wantParams: map[string]string{"chat_id": "100", "message_id": "1"},
		},
		{
			name:   "AnswerInlineQuery",
			update: sandbox.InlineQuery("cats"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerInlineQuery(ctx, []telegram.InlineQueryResult{
					&telegram.InlineQueryResultArticle{ID: "1", Title: "Cat", InputMessageContent: &telegram.InputTextMessageContent{MessageText: "meow"}},
				})
			},
			wantMethod: "answerInlineQuery",
			wantParams: map[string]string{
				"inline_query_id": "inline-1",
				"results":         `[{"input_message_content":{"message_text":"meow"},"id":"1","title":"Cat","type":"article"}]`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			u := gogram.NewUpdate(srv.Client(), tt.update)
			if err := tt.call(t.Context(), u); err != nil {
				t.Fatalf("call error = %v", err)
			}

			reqs := srv.Requests()
			if len(reqs) != 1 || reqs[0].Method != tt.wantMethod {
				t.Fatalf("requests = %v, want one %s", reqs, tt.wantMethod)
			}
			got := map[string]string{}
			for k := range reqs[0].Params {
				got[k] = reqs[0].String(k)
			}
			for k, want := range tt.wantParams {
				if got[k] != want && !jsonEqual(got[k], want) {
					t.Errorf("param %s = %s, want %s", k, got[k], want)
				}
			}
			if len(got) != len(tt.wantParams) {
				t.Errorf("params = %v, want exactly %v", got, tt.wantParams)
			}
		})
	}
}

// jsonEqual reports whether two JSON documents are equal, ignoring key order.
func jsonEqual(a, b string) bool {
	var x, y any
	if json.Unmarshal([]byte(a), &x) != nil || json.Unmarshal([]byte(b), &y) != nil {
		return false
	}
	ja, _ := json.Marshal(x)
	jb, _ := json.Marshal(y)
	return string(ja) == string(jb)
}

func TestUpdateShortcutErrors(t *testing.T) {
	tests := []struct {
		name    string
		update  telegram.Update
		call    func(ctx context.Context, u *gogram.Update) error
		wantErr error
	}{
		{
			name:   "Reply without a chat",
			update: sandbox.InlineQuery("x"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.Reply(ctx, telegram.Text("x"))
				return err
			},
			wantErr: gogram.ErrNoChat,
		},
		{
			name:   "SendChatAction without a chat",
			update: sandbox.InlineQuery("x"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.SendChatAction(ctx, telegram.ChatActionTyping)
			},
			wantErr: gogram.ErrNoChat,
		},
		{
			name:   "BanChatMember without a chat",
			update: sandbox.InlineQuery("x"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.BanChatMember(ctx, 1)
			},
			wantErr: gogram.ErrNoChat,
		},
		{
			name:   "AnswerCallbackQuery on a message",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerCallbackQuery(ctx)
			},
			wantErr: gogram.ErrNoCallbackQuery,
		},
		{
			name:   "EditMessageText on a message",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.EditMessageText(ctx, telegram.Text("x"))
				return err
			},
			wantErr: gogram.ErrNoCallbackQuery,
		},
		{
			name: "EditMessageText without the callback message",
			update: telegram.Update{CallbackQuery: &telegram.CallbackQuery{
				ID: "q", From: sandbox.User, Data: "x",
			}},
			call: func(ctx context.Context, u *gogram.Update) error {
				_, err := u.EditMessageText(ctx, telegram.Text("x"))
				return err
			},
			wantErr: gogram.ErrNoMessage,
		},
		{
			name:   "AnswerInlineQuery on a message",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerInlineQuery(ctx, nil)
			},
			wantErr: gogram.ErrNoInlineQuery,
		},
		{
			name:   "AnswerShippingQuery on a message",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerShippingQuery(ctx, true)
			},
			wantErr: gogram.ErrNoShippingQuery,
		},
		{
			name:   "AnswerPreCheckoutQuery on a message",
			update: sandbox.Message("hi"),
			call: func(ctx context.Context, u *gogram.Update) error {
				return u.AnswerPreCheckoutQuery(ctx, true)
			},
			wantErr: gogram.ErrNoPreCheckoutQuery,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			u := gogram.NewUpdate(srv.Client(), tt.update)
			if err := tt.call(t.Context(), u); !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
			if reqs := srv.Requests(); len(reqs) != 0 {
				t.Errorf("requests = %v, want none", reqs)
			}
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name         string
		opts         []gogram.Option
		getMe        sandbox.HandlerFunc // nil for the default answer
		wantErr      bool
		wantUsername string
		wantCalls    int
	}{
		{
			name:         "looks up the username",
			wantUsername: sandbox.BotUsername,
			wantCalls:    1,
		},
		{
			name:         "username given",
			opts:         []gogram.Option{gogram.WithUsername("@given_bot")},
			wantUsername: "given_bot",
		},
		{
			name: "bad token",
			getMe: func(sandbox.Request) (any, error) {
				return nil, &telegram.Error{Code: 401, Description: "Unauthorized"}
			},
			wantErr:   true,
			wantCalls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			if tt.getMe != nil {
				srv.Handle("getMe", tt.getMe)
			}
			b := gogram.NewWithClient(srv.Client(), append([]gogram.Option{quiet}, tt.opts...)...)

			err := b.Init(t.Context())
			if (err != nil) != tt.wantErr {
				t.Fatalf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := b.Username(); got != tt.wantUsername {
				t.Errorf("Username() = %q, want %q", got, tt.wantUsername)
			}
			if got := len(srv.Calls("getMe")); got != tt.wantCalls {
				t.Errorf("getMe calls = %d, want %d", got, tt.wantCalls)
			}
		})
	}
}

// startPolling runs StartPolling in the background and returns a
// function that stops it and returns its result.
func startPolling(t *testing.T, b *gogram.Bot) (stop func() error) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- b.StartPolling(ctx) }()
	return func() error {
		cancel()
		select {
		case err := <-done:
			return err
		case <-time.After(5 * time.Second):
			t.Fatal("StartPolling did not return after cancel")
			return nil
		}
	}
}

func TestStartPolling(t *testing.T) {
	tests := []struct {
		name     string
		texts    []string
		username bool // whether the bot knows its username
	}{
		{"no updates", nil, true},
		{"one update", []string{"one"}, true},
		{"several updates", []string{"one", "two", "three"}, true},
		{"looks up the username first", []string{"one"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			for _, text := range tt.texts {
				srv.AddUpdates(sandbox.Message(text))
			}
			b := gogram.NewWithClient(srv.Client(), quiet)
			if tt.username {
				b = srv.Bot(quiet)
			}
			b.OnText(func(ctx context.Context, u *gogram.Update) error {
				_, err := u.Reply(ctx, telegram.Text("echo "+u.Message().Text))
				return err
			})

			stop := startPolling(t, b)
			srv.WaitFor("getUpdates", 1)
			replies := srv.WaitFor("sendMessage", len(tt.texts))

			var got, want []string
			for _, r := range replies {
				got = append(got, r.String("text"))
			}
			for _, text := range tt.texts {
				want = append(want, "echo "+text)
			}
			if !slices.Equal(got, want) { // one chat, so in order
				t.Errorf("replies = %q, want %q", got, want)
			}

			if err := b.StartPolling(t.Context()); !errors.Is(err, gogram.ErrAlreadyPolling) {
				t.Errorf("second StartPolling() error = %v, want ErrAlreadyPolling", err)
			}
			if err := stop(); err != nil {
				t.Fatalf("StartPolling() = %v, want nil after cancel", err)
			}

			polls := srv.Calls("getUpdates")
			if first := polls[0].Int("offset"); first != 0 {
				t.Errorf("first offset = %d, want 0", first)
			}
			wantLast := int64(0)
			if len(tt.texts) > 0 {
				wantLast = int64(len(tt.texts) + 1)
			}
			if last := polls[len(polls)-1].Int("offset"); last != wantLast {
				t.Errorf("final offset = %d, want %d", last, wantLast)
			}
			if got, want := len(srv.Calls("getMe")), map[bool]int{true: 0, false: 1}[tt.username]; got != want {
				t.Errorf("getMe calls = %d, want %d", got, want)
			}
		})
	}
}

// fromChat returns a text message in the given chat.
func fromChat(chatID int64, text string) telegram.Update {
	u := sandbox.Message(text)
	u.Message.Chat.ID = chatID
	return u
}

func TestPollingConcurrency(t *testing.T) {
	tests := []struct {
		name        string
		concurrency int
		updates     []telegram.Update
		// slow makes the handler of the first update slow, so that later
		// updates would overtake it if they could.
		slow bool
		// together requires all updates to be running at the same time.
		together    bool
		wantMaxBusy int32 // upper bound on handlers running at once; 0 for none
	}{
		{
			name:        "a chat's updates run in order",
			updates:     []telegram.Update{fromChat(1, "a"), fromChat(1, "b"), fromChat(1, "c"), fromChat(1, "d")},
			slow:        true,
			wantMaxBusy: 1,
		},
		{
			name:     "different chats run concurrently",
			updates:  []telegram.Update{fromChat(1, "a"), fromChat(2, "b"), fromChat(3, "c")},
			together: true,
		},
		{
			name:        "concurrency limit",
			concurrency: 2,
			updates:     []telegram.Update{fromChat(1, "a"), fromChat(2, "b"), fromChat(3, "c"), fromChat(4, "d"), fromChat(5, "e")},
			wantMaxBusy: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			srv.AddUpdates(tt.updates...)

			var (
				mu         sync.Mutex
				order      = map[int64][]string{}
				errs       []error
				busy       atomic.Int32
				maxBusy    atomic.Int32
				started    sync.WaitGroup
				handled    sync.WaitGroup
				allStarted = make(chan struct{})
			)
			started.Add(len(tt.updates))
			handled.Add(len(tt.updates))
			go func() { started.Wait(); close(allStarted) }()

			b := srv.Bot(quiet, gogram.WithConcurrency(tt.concurrency),
				gogram.WithErrorHandler(func(_ context.Context, _ *gogram.Update, err error) {
					mu.Lock()
					errs = append(errs, err)
					mu.Unlock()
				}))
			b.OnText(func(ctx context.Context, u *gogram.Update) error {
				defer handled.Done()
				n := busy.Add(1)
				defer busy.Add(-1)
				for m := maxBusy.Load(); n > m && !maxBusy.CompareAndSwap(m, n); m = maxBusy.Load() {
				}
				started.Done()
				switch {
				case tt.together:
					select {
					case <-allStarted:
					case <-time.After(2 * time.Second):
						return errors.New("updates of different chats did not run together")
					}
				case tt.slow && u.Message().Text == "a":
					time.Sleep(50 * time.Millisecond)
				default:
					time.Sleep(10 * time.Millisecond)
				}
				mu.Lock()
				order[u.ChatID()] = append(order[u.ChatID()], u.Message().Text)
				mu.Unlock()
				return nil
			})

			stop := startPolling(t, b)
			handled.Wait()
			if err := stop(); err != nil {
				t.Fatalf("StartPolling() = %v", err)
			}

			mu.Lock()
			defer mu.Unlock()
			for _, err := range errs {
				t.Error(err)
			}
			want := map[int64][]string{}
			for _, u := range tt.updates {
				want[u.Message.Chat.ID] = append(want[u.Message.Chat.ID], u.Message.Text)
			}
			for chat, texts := range want {
				if !slices.Equal(order[chat], texts) {
					t.Errorf("chat %d handled %q, want %q in order", chat, order[chat], texts)
				}
			}
			if tt.wantMaxBusy > 0 && maxBusy.Load() > tt.wantMaxBusy {
				t.Errorf("%d handlers ran at once, want at most %d", maxBusy.Load(), tt.wantMaxBusy)
			}
		})
	}
}

func TestShutdown(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
		stuck   bool // the handler runs until its context is cancelled
		wantErr error
	}{
		{"waits for running handlers", 0, false, nil},
		{"handlers finishing in time", time.Second, false, nil},
		{"stuck handler is cancelled after the timeout", 50 * time.Millisecond, true, gogram.ErrShutdownTimeout},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			srv.AddUpdates(sandbox.Message("work"))
			b := srv.Bot(quiet, gogram.WithShutdownTimeout(tt.timeout))

			running := make(chan struct{})
			finished := make(chan error, 1)
			b.OnText(func(ctx context.Context, u *gogram.Update) error {
				close(running)
				if tt.stuck {
					<-ctx.Done()
				} else {
					time.Sleep(50 * time.Millisecond)
				}
				// Replies still go out during shutdown, unless cancelled.
				_, err := u.Reply(ctx, telegram.Text("done"))
				finished <- err
				return nil
			})

			stop := startPolling(t, b)
			<-running
			if err := stop(); !errors.Is(err, tt.wantErr) {
				t.Fatalf("StartPolling() = %v, want %v", err, tt.wantErr)
			}
			select {
			case err := <-finished:
				if tt.stuck && !errors.Is(err, context.Canceled) {
					t.Errorf("reply error = %v, want context.Canceled", err)
				}
				if !tt.stuck && err != nil {
					t.Errorf("reply error = %v, want the reply sent during shutdown", err)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("handler did not finish")
			}
		})
	}
}

func TestWebhookHandler(t *testing.T) {
	update, _ := json.Marshal(sandbox.Message("via webhook"))
	tooLarge := fmt.Sprintf(`{"update_id":1,"message":{"text":%q}}`, strings.Repeat("a", 1<<20))

	tests := []struct {
		name        string
		secret      string // configured on the handler
		method      string
		header      string // sent in X-Telegram-Bot-Api-Secret-Token
		body        string
		wantCode    int
		wantHandled bool
	}{
		{"valid", "s3cret", http.MethodPost, "s3cret", string(update), http.StatusOK, true},
		{"no secret configured", "", http.MethodPost, "", string(update), http.StatusOK, true},
		{"wrong secret", "s3cret", http.MethodPost, "wrong", string(update), http.StatusForbidden, false},
		{"missing secret", "s3cret", http.MethodPost, "", string(update), http.StatusForbidden, false},
		{"not a POST", "s3cret", http.MethodGet, "s3cret", "", http.StatusMethodNotAllowed, false},
		{"invalid JSON", "s3cret", http.MethodPost, "s3cret", "{", http.StatusBadRequest, false},
		{"body too large", "s3cret", http.MethodPost, "s3cret", tooLarge, http.StatusBadRequest, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			b := gogram.NewWithClient(srv.Client(), quiet)
			var handled string
			b.OnText(func(ctx context.Context, u *gogram.Update) error {
				handled = u.Message().Text
				return nil
			})

			req := httptest.NewRequest(tt.method, "/bot", strings.NewReader(tt.body))
			if tt.header != "" {
				req.Header.Set("X-Telegram-Bot-Api-Secret-Token", tt.header)
			}
			rec := httptest.NewRecorder()
			b.WebhookHandler(tt.secret).ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantCode)
			}
			// The update must be handled before the response is written.
			if got := handled == "via webhook"; got != tt.wantHandled {
				t.Errorf("handled = %v (%q), want %v", got, handled, tt.wantHandled)
			}
			// The first handled update looks up the bot's username.
			if wantGetMe := map[bool]int{true: 1, false: 0}[tt.wantHandled]; len(srv.Calls("getMe")) != wantGetMe {
				t.Errorf("getMe calls = %d, want %d", len(srv.Calls("getMe")), wantGetMe)
			}
		})
	}
}
