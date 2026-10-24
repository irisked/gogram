package telegram_test

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"

	"github.com/irisked/gogram/telegram"
)

// command returns a message whose text starts with a command entity of
// the given length, as Telegram sends it.
func command(text string, length int) *telegram.Message {
	return &telegram.Message{
		Text:     text,
		Entities: []telegram.MessageEntity{{Type: "bot_command", Offset: 0, Length: length}},
	}
}

func TestMessageCommand(t *testing.T) {
	tests := []struct {
		name          string
		msg           *telegram.Message
		wantIsCommand bool
		wantCommand   string
		wantWithAt    string
		wantBot       string
		wantArgs      string
	}{
		{
			name:          "bare command",
			msg:           command("/help", 5),
			wantIsCommand: true, wantCommand: "help", wantWithAt: "help",
		},
		{
			name:          "command with arguments",
			msg:           command("/ban 42 spam", 4),
			wantIsCommand: true, wantCommand: "ban", wantWithAt: "ban", wantArgs: "42 spam",
		},
		{
			name:          "command with bot name",
			msg:           command("/start@my_bot  payload here", len("/start@my_bot")),
			wantIsCommand: true, wantCommand: "start", wantWithAt: "start@my_bot", wantBot: "my_bot", wantArgs: "payload here",
		},
		{
			name: "entity longer than the text",
			msg:  command("/st", 6),
		},
		{
			name: "plain text",
			msg:  &telegram.Message{Text: "/start without an entity"},
		},
		{
			name: "command entity not at the start",
			msg: &telegram.Message{
				Text:     "hi /start",
				Entities: []telegram.MessageEntity{{Type: "bot_command", Offset: 3, Length: 6}},
			},
		},
		{
			name: "other entity at the start",
			msg: &telegram.Message{
				Text:     "@user hi",
				Entities: []telegram.MessageEntity{{Type: "mention", Offset: 0, Length: 5}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msg.IsCommand(); got != tt.wantIsCommand {
				t.Errorf("IsCommand() = %v, want %v", got, tt.wantIsCommand)
			}
			if got := tt.msg.Command(); got != tt.wantCommand {
				t.Errorf("Command() = %q, want %q", got, tt.wantCommand)
			}
			if got := tt.msg.CommandWithAt(); got != tt.wantWithAt {
				t.Errorf("CommandWithAt() = %q, want %q", got, tt.wantWithAt)
			}
			if got := tt.msg.CommandBot(); got != tt.wantBot {
				t.Errorf("CommandBot() = %q, want %q", got, tt.wantBot)
			}
			if got := tt.msg.CommandArguments(); got != tt.wantArgs {
				t.Errorf("CommandArguments() = %q, want %q", got, tt.wantArgs)
			}
		})
	}
}

func TestMessageType(t *testing.T) {
	loc := &telegram.Location{}
	tests := []struct {
		name string
		msg  *telegram.Message
		want telegram.MessageType
	}{
		{"text", &telegram.Message{Text: "hi"}, telegram.MessageTypeText},
		{"command", command("/start", 6), telegram.MessageTypeCommand},
		{"photo", &telegram.Message{Photo: []telegram.PhotoSize{{}}}, telegram.MessageTypePhoto},
		{"video", &telegram.Message{Video: &telegram.Video{}}, telegram.MessageTypeVideo},
		{"animation wins over document", &telegram.Message{Animation: &telegram.Animation{}, Document: &telegram.Document{}}, telegram.MessageTypeAnimation},
		{"document", &telegram.Message{Document: &telegram.Document{}}, telegram.MessageTypeDocument},
		{"venue wins over location", &telegram.Message{Venue: &telegram.Venue{}, Location: loc}, telegram.MessageTypeVenue},
		{"location", &telegram.Message{Location: loc}, telegram.MessageTypeLocation},
		{"new members", &telegram.Message{NewChatMembers: []telegram.User{{}}}, telegram.MessageTypeNewChatMembers},
		{"deleted photo", &telegram.Message{DeleteChatPhoto: true}, telegram.MessageTypeDeleteChatPhoto},
		{"other", &telegram.Message{}, telegram.MessageTypeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msg.Type(); got != tt.want {
				t.Errorf("Type() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestUpdateType(t *testing.T) {
	tests := []struct {
		name string
		json string
		want telegram.UpdateType
	}{
		{"message", `{"update_id":1,"message":{"message_id":1}}`, telegram.UpdateMessage},
		{"edited message", `{"update_id":1,"edited_message":{"message_id":1}}`, telegram.UpdateEditedMessage},
		{"channel post", `{"update_id":1,"channel_post":{"message_id":1}}`, telegram.UpdateChannelPost},
		{"callback query", `{"update_id":1,"callback_query":{"id":"q"}}`, telegram.UpdateCallbackQuery},
		{"inline query", `{"update_id":1,"inline_query":{"id":"q"}}`, telegram.UpdateInlineQuery},
		{"poll", `{"update_id":1,"poll":{"id":"p"}}`, telegram.UpdatePoll},
		{"my chat member", `{"update_id":1,"my_chat_member":{}}`, telegram.UpdateMyChatMember},
		{"poll answer", `{"update_id":1,"poll_answer":{"poll_id":"p"}}`, telegram.UpdatePollAnswer},
		{"chat join request", `{"update_id":1,"chat_join_request":{}}`, telegram.UpdateChatJoinRequest},
		{"message reaction", `{"update_id":1,"message_reaction":{}}`, telegram.UpdateMessageReaction},
		{"business message", `{"update_id":1,"business_message":{"message_id":1}}`, telegram.UpdateBusinessMessage},
		{"type the package doesn't know", `{"update_id":1,"future_update":{}}`, ""},
		{"no payload", `{"update_id":1}`, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u telegram.Update
			if err := json.Unmarshal([]byte(tt.json), &u); err != nil {
				t.Fatal(err)
			}
			if got := u.Type(); got != tt.want {
				t.Errorf("Type() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCallbackButton(t *testing.T) {
	tests := []struct {
		name      string
		route     string
		args      []string
		wantData  string
		wantRoute string
		wantArgs  []string
	}{
		{"route only", "noop", nil, "noop", "noop", nil},
		{"one argument", "vote", []string{"up"}, "vote:up", "vote", []string{"up"}},
		{"several arguments", "like", []string{"42", "7"}, "like:42,7", "like", []string{"42", "7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := telegram.CallbackButton("text", tt.route, tt.args...)
			if b.CallbackData != tt.wantData {
				t.Errorf("CallbackData = %q, want %q", b.CallbackData, tt.wantData)
			}

			// The data must round-trip through the CallbackQuery Telegram sends back.
			cq := telegram.CallbackQuery{Data: b.CallbackData}
			if got := cq.Route(); got != tt.wantRoute {
				t.Errorf("Route() = %q, want %q", got, tt.wantRoute)
			}
			if got := cq.Args(); !slices.Equal(got, tt.wantArgs) {
				t.Errorf("Args() = %q, want %q", got, tt.wantArgs)
			}
		})
	}
}

func TestCallbackQueryRouteAndArgs(t *testing.T) {
	tests := []struct {
		data      string
		wantRoute string
		wantArgs  []string
	}{
		{"", "", nil},
		{"vote", "vote", nil},
		{"vote:", "vote", nil},
		{"vote:up", "vote", []string{"up"}},
		{"vote:up,", "vote", []string{"up", ""}},
		{"page:2:extra", "page", []string{"2:extra"}},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			cq := telegram.CallbackQuery{Data: tt.data}
			if got := cq.Route(); got != tt.wantRoute {
				t.Errorf("Route() = %q, want %q", got, tt.wantRoute)
			}
			if got := cq.Args(); !slices.Equal(got, tt.wantArgs) {
				t.Errorf("Args() = %q, want %q", got, tt.wantArgs)
			}
		})
	}
}

func TestCallbackQueryTarget(t *testing.T) {
	tests := []struct {
		name   string
		query  telegram.CallbackQuery
		want   telegram.MessageTarget
		wantOK bool
	}{
		{
			name:   "inline message",
			query:  telegram.CallbackQuery{InlineMessageID: "abc"},
			want:   telegram.InlineMessage("abc"),
			wantOK: true,
		},
		{
			name:   "chat message",
			query:  telegram.CallbackQuery{Message: &telegram.Message{MessageID: 3, Chat: &telegram.Chat{ID: 9}}},
			want:   telegram.ChatMessage(9, 3),
			wantOK: true,
		},
		{
			name:  "message without chat",
			query: telegram.CallbackQuery{Message: &telegram.Message{MessageID: 3}},
		},
		{
			name: "no message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.query.Target()
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("Target() = %+v, %v; want %+v, %v", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestReplyMarkupJSON(t *testing.T) {
	resized := telegram.NewReplyKeyboard(telegram.ContactButton("Share phone")).Row(telegram.TextButton("Cancel"))
	resized.ResizeKeyboard = true

	tests := []struct {
		name   string
		markup telegram.ReplyMarkup
		want   string
	}{
		{
			name: "inline keyboard",
			markup: telegram.NewInlineKeyboard(
				telegram.CallbackButton("Yes", "vote", "yes"),
				telegram.URLButton("Rules", "https://example.com"),
			),
			want: `{"inline_keyboard":[[{"text":"Yes","callback_data":"vote:yes"},{"text":"Rules","url":"https://example.com"}]]}`,
		},
		{
			name:   "inline keyboard with rows",
			markup: telegram.NewInlineKeyboard(telegram.CallbackButton("A", "a")).Row(telegram.CallbackButton("B", "b")),
			want:   `{"inline_keyboard":[[{"text":"A","callback_data":"a"}],[{"text":"B","callback_data":"b"}]]}`,
		},
		{
			name:   "reply keyboard",
			markup: resized,
			want:   `{"keyboard":[[{"text":"Share phone","request_contact":true}],[{"text":"Cancel"}]],"resize_keyboard":true}`,
		},
		{
			name:   "remove keyboard",
			markup: telegram.RemoveKeyboard(),
			want:   `{"remove_keyboard":true}`,
		},
		{
			name:   "force reply",
			markup: telegram.NewForceReply(),
			want:   `{"force_reply":true}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.markup)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tt.want {
				t.Errorf("JSON = %s\nwant   %s", got, tt.want)
			}
		})
	}
}

func TestChatIDJSON(t *testing.T) {
	tests := []struct {
		name string
		id   telegram.ChatID
		json string
	}{
		{"numeric", telegram.ChatID{ID: -1001234}, `-1001234`},
		{"username", telegram.ChatUsername("mychannel"), `"@mychannel"`},
		{"username with @", telegram.ChatUsername("@mychannel"), `"@mychannel"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.id)
			if err != nil || string(b) != tt.json {
				t.Fatalf("Marshal = %s, %v; want %s", b, err, tt.json)
			}
			var back telegram.ChatID
			if err := json.Unmarshal(b, &back); err != nil || back != tt.id {
				t.Errorf("Unmarshal(%s) = %+v, %v; want %+v", b, back, err, tt.id)
			}
		})
	}
}

// TestNestedInterfaceDecoding decodes updates whose objects hold
// interface-typed fields.
func TestNestedInterfaceDecoding(t *testing.T) {
	tests := []struct {
		name  string
		json  string
		check func(t *testing.T, u *telegram.Update)
	}{
		{
			name: "chat member status change",
			json: `{"update_id":1,"my_chat_member":{"chat":{"id":-5,"type":"group"},"from":{"id":2},"date":1,
				"old_chat_member":{"status":"left","user":{"id":9}},
				"new_chat_member":{"status":"kicked","user":{"id":9},"until_date":99}}}`,
			check: func(t *testing.T, u *telegram.Update) {
				if _, ok := u.MyChatMember.OldChatMember.(*telegram.ChatMemberLeft); !ok {
					t.Errorf("OldChatMember = %T, want *ChatMemberLeft", u.MyChatMember.OldChatMember)
				}
				banned, ok := u.MyChatMember.NewChatMember.(*telegram.ChatMemberBanned)
				if !ok || banned.UntilDate != 99 || banned.User.ID != 9 {
					t.Errorf("NewChatMember = %#v, want *ChatMemberBanned until 99", u.MyChatMember.NewChatMember)
				}
			},
		},
		{
			name: "forwarded message",
			json: `{"update_id":1,"message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},
				"forward_origin":{"type":"user","date":5,"sender_user":{"id":3,"first_name":"Bo"}}}}`,
			check: func(t *testing.T, u *telegram.Update) {
				origin, ok := u.Message.ForwardOrigin.(*telegram.MessageOriginUser)
				if !ok || origin.SenderUser.ID != 3 {
					t.Errorf("ForwardOrigin = %#v, want *MessageOriginUser from 3", u.Message.ForwardOrigin)
				}
			},
		},
		{
			name: "reactions",
			json: `{"update_id":1,"message_reaction":{"chat":{"id":1,"type":"private"},"message_id":4,"date":1,
				"old_reaction":[],"new_reaction":[{"type":"emoji","emoji":"🔥"},{"type":"paid"}]}}`,
			check: func(t *testing.T, u *telegram.Update) {
				want := []telegram.ReactionType{&telegram.ReactionTypeEmoji{Emoji: "🔥"}, &telegram.ReactionTypePaid{}}
				if got := u.MessageReaction.NewReaction; !reflect.DeepEqual(got, want) {
					t.Errorf("NewReaction = %#v, want %#v", got, want)
				}
			},
		},
		{
			name: "callback on an inaccessible message",
			json: `{"update_id":1,"callback_query":{"id":"q","from":{"id":2},"chat_instance":"c","data":"x",
				"message":{"chat":{"id":7,"type":"private"},"message_id":3,"date":0}}}`,
			check: func(t *testing.T, u *telegram.Update) {
				m := u.CallbackQuery.Message
				if m == nil || m.IsAccessible() || m.Chat.ID != 7 || m.MessageID != 3 {
					t.Errorf("Message = %#v, want an inaccessible message 3 in chat 7", m)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u telegram.Update
			if err := json.Unmarshal([]byte(tt.json), &u); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			tt.check(t, &u)

			// Encoding and decoding again keeps the values.
			b, err := json.Marshal(&u)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			var again telegram.Update
			if err := json.Unmarshal(b, &again); err != nil {
				t.Fatalf("Unmarshal re-encoded: %v", err)
			}
			if !reflect.DeepEqual(again, u) {
				t.Errorf("round trip changed the update:\n got %#v\nwant %#v", again, u)
			}
		})
	}
}
