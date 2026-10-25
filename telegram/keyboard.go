package telegram

import (
	"fmt"
	"strings"
)

// ReplyMarkup is the reply_markup of a message: an *InlineKeyboardMarkup,
// a *ReplyKeyboardMarkup, a *ReplyKeyboardRemove or a *ForceReply.
type ReplyMarkup interface {
	replyMarkup()
}

func (*InlineKeyboardMarkup) replyMarkup() {}
func (*ReplyKeyboardMarkup) replyMarkup()  {}
func (*ReplyKeyboardRemove) replyMarkup()  {}
func (*ForceReply) replyMarkup()           {}

// NewInlineKeyboard creates an inline keyboard whose first row holds
// buttons. Add more rows with Row.
//
//	kb := telegram.NewInlineKeyboard(
//		telegram.CallbackButton("Yes", "vote", "yes"),
//		telegram.CallbackButton("No", "vote", "no"),
//	).Row(
//		telegram.URLButton("Docs", "https://core.telegram.org/bots/api"),
//	)
func NewInlineKeyboard(buttons ...InlineKeyboardButton) *InlineKeyboardMarkup {
	k := &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{}}
	if len(buttons) > 0 {
		k.Row(buttons...)
	}
	return k
}

// Row appends a row of buttons and returns k for chaining.
func (k *InlineKeyboardMarkup) Row(buttons ...InlineKeyboardButton) *InlineKeyboardMarkup {
	k.InlineKeyboard = append(k.InlineKeyboard, buttons)
	return k
}

// MaxCallbackData is the maximum size of callback data, in bytes.
const MaxCallbackData = 64

// CallbackButton sends a callback query when pressed. The callback data
// is route, followed by ":" and the comma-separated args when there are
// any, so that gogram's Bot.OnCallback(route, ...) handles it:
//
//	telegram.CallbackButton("Like", "like", postID) // data "like:42"
//
// Read the arguments back with CallbackQuery.Args. The data must fit in
// MaxCallbackData bytes; requests with a longer one fail before they are
// sent.
func CallbackButton(text, route string, args ...string) InlineKeyboardButton {
	data := route
	if len(args) > 0 {
		data += ":" + strings.Join(args, ",")
	}
	return InlineKeyboardButton{Text: text, CallbackData: data}
}

// validate is called by the request encoder.
func (b InlineKeyboardButton) validate() error {
	if len(b.CallbackData) > MaxCallbackData {
		return fmt.Errorf("button %q: callback data is %d bytes, more than the %d Telegram allows",
			b.Text, len(b.CallbackData), MaxCallbackData)
	}
	return nil
}

// URLButton opens url when pressed.
func URLButton(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, URL: url}
}

// WebAppButton opens a Web App when pressed.
func WebAppButton(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, WebApp: &WebAppInfo{URL: url}}
}

// SwitchInlineQueryButton lets the user pick a chat and starts an inline
// query there with query pre-filled.
func SwitchInlineQueryButton(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, SwitchInlineQuery: &query}
}

// SwitchInlineQueryCurrentChatButton starts an inline query in the
// current chat with query pre-filled.
func SwitchInlineQueryCurrentChatButton(text, query string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: &query}
}

// GameButton launches the game of a SendGame message. It must be the
// first button of the first row.
func GameButton(text string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, CallbackGame: &CallbackGame{}}
}

// PayButton pays an invoice. It must be the first button of the first row.
func PayButton(text string) InlineKeyboardButton {
	return InlineKeyboardButton{Text: text, Pay: true}
}

// NewReplyKeyboard creates a reply keyboard whose first row holds
// buttons. Add more rows with Row, and set ResizeKeyboard and the other
// fields directly.
func NewReplyKeyboard(buttons ...KeyboardButton) *ReplyKeyboardMarkup {
	k := &ReplyKeyboardMarkup{Keyboard: [][]KeyboardButton{}}
	if len(buttons) > 0 {
		k.Row(buttons...)
	}
	return k
}

// Row appends a row of buttons and returns k for chaining.
func (k *ReplyKeyboardMarkup) Row(buttons ...KeyboardButton) *ReplyKeyboardMarkup {
	k.Keyboard = append(k.Keyboard, buttons)
	return k
}

// TextButton sends its text when pressed.
func TextButton(text string) KeyboardButton {
	return KeyboardButton{Text: text}
}

// ContactButton shares the user's phone number when pressed.
func ContactButton(text string) KeyboardButton {
	return KeyboardButton{Text: text, RequestContact: true}
}

// LocationButton shares the user's location when pressed.
func LocationButton(text string) KeyboardButton {
	return KeyboardButton{Text: text, RequestLocation: true}
}

// RemoveKeyboard returns markup that removes the custom keyboard.
func RemoveKeyboard() *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{RemoveKeyboard: true}
}

// NewForceReply returns markup that forces a reply.
func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}
