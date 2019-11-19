package method

import (
	"github.com/irisked/gogram/args"
	"github.com/irisked/gogram/telegram/keyboard"
)

// GetGameHighScores contains information about getGameHighScores telegram api method.
type GetGameHighScores struct {
	UserID          int    `json:"user_id"`
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	SimpleMethod
}

// NewGetGameHighScores creates GetGameHighScores.
func NewGetGameHighScores(userID int, message args.MessageID) *GetGameHighScores {
	method := new(GetGameHighScores)
	method.UserID = userID
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	return method
}

// Method returns telegram api method endpoint.
func (m *GetGameHighScores) Method() string {
	return "getGameHighScores"
}

// SetGameScores contains information about SetGameScore telegram api method.
type SetGameScores struct {
	UserID             int    `json:"user_id"`
	Score              int    `json:"score"`
	ChatID             int64  `json:"chat_id,omitempty"`
	MessageID          int    `json:"message_id,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	SimpleMethod
}

// SetGameScoresOption interface for optional parameters.
type SetGameScoresOption interface {
	ApplySetGameScoresOption(*SetGameScores)
}

// NewSetGameScores creates SetGameScores.
func NewSetGameScores(userID, score int, message args.MessageID, options ...SetGameScoresOption) *SetGameScores {
	method := new(SetGameScores)
	method.UserID = userID
	method.Score = score
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	for _, opt := range options {
		opt.ApplySetGameScoresOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (m *SetGameScores) Method() string {
	return "setGameScore"
}

// SendGame contains information about sendGame telegram api method.
type SendGame struct {
	ChatID              int               `json:"chat_id"`
	GameShortName       string            `json:"game_short_name"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// SendGameOption interface for optional parameters.
type SendGameOption interface {
	ApplySendGameOption(*SendGame)
}

// NewSendGame creates SendGame.
func NewSendGame(chatID int, name string, options ...SendGameOption) *SendGame {
	method := new(SendGame)
	method.ChatID = chatID
	method.GameShortName = name
	for _, opt := range options {
		opt.ApplySendGameOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendGame) Method() string {
	return "sendGame"
}
