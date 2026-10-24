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

// NewSetGameScores creates SetGameScores.
func NewSetGameScores(userID, score int, message args.MessageID, fn ...func(*SetGameScores)) *SetGameScores {
	method := new(SetGameScores)
	method.UserID = userID
	method.Score = score
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	for _, f := range fn {
		f(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (m *SetGameScores) Method() string {
	return "setGameScore"
}

// SetForce sets force
func (m *SetGameScores) SetForce(v bool) {
	m.Force = v
}

// SetDisableEditMessage sets disable_edit_message
func (m *SetGameScores) SetDisableEditMessage(v bool) {
	m.DisableEditMessage = v
}

// SendGame contains information about sendGame telegram api method.
type SendGame struct {
	ChatID              int               `json:"chat_id"`
	GameShortName       string            `json:"game_short_name"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty"`
	SimpleMethod
}

// NewSendGame creates SendGame.
func NewSendGame(chatID int, name string, fn ...func(*SendGame)) *SendGame {
	method := new(SendGame)
	method.ChatID = chatID
	method.GameShortName = name
	for _, f := range fn {
		f(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendGame) Method() string {
	return "sendGame"
}

// SetDisableNotification sets disable_notification
func (r *SendGame) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (r *SendGame) SetReplyToMessageID(id int) {
	r.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (r *SendGame) SetReplyMarkup(kb keyboard.Markup) {
	r.ReplyMarkup = kb
}
