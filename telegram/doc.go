// Package telegram is a client for the Telegram Bot API.
//
// It can be used on its own or through the gogram bot framework. The
// methods, their requests and the API types are generated from the Bot
// API reference, so every method and type of the supported Bot API
// version is available.
//
// Each API method is a Client method that takes its required parameters
// as arguments and its optional parameters as options:
//
//	client := telegram.New(token)
//	msg, err := client.SendMessage(ctx, chatID, telegram.Markdown("*Hello*"),
//		telegram.WithReplyMarkup(keyboard),
//		telegram.WithDisableNotification(),
//	)
//
// # Options
//
// The option for a parameter is named after it: WithReplyMarkup sets
// reply_markup. One option works with every method that has the
// parameter, and using it with a method that doesn't is a compile error.
// Each method's documentation links the options it accepts.
//
// Formatted text (text, caption and the like) is a TextData: Text,
// Markdown, MarkdownV2 or HTML, which sets the parse mode along with the
// text.
//
// # Files
//
// Methods that send files take an *InputFile: FileID or FileURL to reuse
// a file Telegram can already reach, or LocalPath, FileReader and
// LocalFile to upload new content. Uploads switch the request to
// multipart/form-data automatically.
//
// # Requests
//
// Every method also has a request type (SendMessage for sendMessage)
// with a constructor (NewSendMessage). Client.Do sends any Request,
// including ones you define for methods this package doesn't cover, and
// requests addressed to a chat by username (see ChatID).
//
// # Interfaces
//
// Objects that come in several kinds, such as ChatMember or InputMedia,
// are interfaces with one struct per kind; see UnknownVariant.
//
// # Errors
//
// Calls rejected by Telegram return an *Error carrying the error code,
// description and, for rate limits, how long to wait.
package telegram

//go:generate go run ../internal/cmd/genapi -spec ../internal/apispec/api.html -telegram . -gogram ..
