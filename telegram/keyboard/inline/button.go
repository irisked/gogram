package inline

import (
	"strings"
)

// Button represents one button of an inline keyboard
// You must use exactly one of the optional fields
type Button struct {
	Text                         string `json:"text"`
	URL                          string `json:"url,omitempty"`                              // Optional
	CallbackData                 string `json:"callback_data,omitempty"`                    // Optional
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`              // Optional
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"` // Optional
	CallbackGame                 string `json:"callback_game,omitempty"`                    // Optional
	Pay                          bool   `json:"pay,omitempty"`                              // Optional
}

// NewButton creates a button
func NewButton(text string, option func(*Button)) Button {
	button := new(Button)
	button.Text = text
	option(button)
	return *button
}

// URL sets URL and defaults everything other.
func URL(url string) func(*Button) {
	return func(btn *Button) {
		btn.URL = url
	}
}

// Callback sets CallbackData and defaults everything other.
func Callback(data string, params ...string) func(*Button) {
	return func(btn *Button) {
		var builder strings.Builder
		builder.WriteString(data)
		builder.WriteString(":")
		length := len(params)
		isLast := func (index, length int) bool {
			return index == length
		}
		for index, element := range params {
			builder.WriteString(element)
			if !isLast(index, length - 1) {
				builder.WriteString(",")
			}
		}
		btn.CallbackData = builder.String()
	}
}

// SwitchInlineQuery sets SwitchInlineQuery and defaults everything other.
func SwitchInlineQuery(query string) func(*Button) {
	return func(btn *Button) {
		btn.SwitchInlineQuery = query
		btn.URL = ""
		btn.CallbackData = ""
		btn.SwitchInlineQueryCurrentChat = ""
		btn.CallbackGame = ""
		btn.Pay = false
	}
}

// SwitchInlineQueryCurrentChat sets SwitchInlineQueryCurrentChat and defaults everything other.
func SwitchInlineQueryCurrentChat(query string) func(*Button) {
	return func(btn *Button) {
		btn.SwitchInlineQueryCurrentChat = query
	}
}

// CallbackGame sets CallbackGame and defaults everything other.
func CallbackGame(cbgame string) func(*Button) {
	return func(btn *Button) {
		btn.CallbackGame = cbgame
	}
}

// Pay sets Pay and defaults everything other.
func Pay() func(*Button) {
	return func(btn *Button) {
		btn.Pay = true
	}
}
