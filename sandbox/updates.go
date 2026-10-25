package sandbox

import (
	"strings"
	"time"

	"github.com/irisked/gogram/telegram"
)

// The users and chats of the updates built by this package.
const (
	BotUsername       = "test_bot"
	UserID      int64 = 100
	// ChatID is the private chat with User; private chats have the ID of
	// their user.
	ChatID  int64 = UserID
	GroupID int64 = -1001000000001
)

var (
	// Bot is the test bot, as returned by getMe.
	Bot = &telegram.User{ID: 1, IsBot: true, FirstName: "Test Bot", Username: BotUsername}
	// User sends the updates built by this package.
	User = &telegram.User{ID: UserID, FirstName: "Test", Username: "test_user", LanguageCode: "en"}
)

func privateChat() *telegram.Chat {
	return &telegram.Chat{ID: ChatID, Type: "private", FirstName: User.FirstName, Username: User.Username}
}

func groupChat() *telegram.Chat {
	return &telegram.Chat{ID: GroupID, Type: "supergroup", Title: "Test Group"}
}

func message(chat *telegram.Chat, text string) *telegram.Message {
	m := &telegram.Message{MessageID: 1, Date: time.Now().Unix(), Chat: chat, From: User, Text: text}
	if strings.HasPrefix(text, "/") {
		command, _, _ := strings.Cut(text, " ")
		m.Entities = []telegram.MessageEntity{{Type: "bot_command", Offset: 0, Length: len(command)}}
	}
	return m
}

// Message returns a text message from User in the private chat. Text
// starting with "/" is marked as a command, as Telegram does.
func Message(text string) telegram.Update {
	return telegram.Update{Message: message(privateChat(), text)}
}

// Command returns a command message, such as "/start" or
// "/ban 42", from User in the private chat.
func Command(text string) telegram.Update {
	if !strings.HasPrefix(text, "/") {
		text = "/" + text
	}
	return Message(text)
}

// GroupMessage returns a text message from User in the group GroupID.
func GroupMessage(text string) telegram.Update {
	return telegram.Update{Message: message(groupChat(), text)}
}

// Callback returns a press by User of a button with the given callback
// data, attached to a bot message in the private chat.
func Callback(data string) telegram.Update {
	return telegram.Update{CallbackQuery: &telegram.CallbackQuery{
		ID:           "callback-1",
		From:         User,
		Message:      &telegram.Message{MessageID: 1, Date: time.Now().Unix(), Chat: privateChat(), From: Bot},
		ChatInstance: "chat-instance-1",
		Data:         data,
	}}
}

// InlineQuery returns an inline query by User.
func InlineQuery(query string) telegram.Update {
	return telegram.Update{InlineQuery: &telegram.InlineQuery{ID: "inline-1", From: User, Query: query}}
}
