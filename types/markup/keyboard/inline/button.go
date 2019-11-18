package inline

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
func NewButton(text string, option func(*Button)) *Button {
	button := new(Button)
	button.Text = text
	option(button)
	return button
}

// WithURL sets URL and defaults everything other.
func WithURL(url string) func(*Button) {
	return func(btn *Button) {
		btn.URL = url
		btn.CallbackData = ""
		btn.SwitchInlineQuery = ""
		btn.SwitchInlineQueryCurrentChat = ""
		btn.CallbackGame = ""
		btn.Pay = false
	}
}

// WithCallbackData sets CallbackData and defaults everything other.
func WithCallbackData(data string) func(*Button) {
	return func(btn *Button) {
		btn.CallbackData = data
		btn.URL = ""
		btn.SwitchInlineQuery = ""
		btn.SwitchInlineQueryCurrentChat = ""
		btn.CallbackGame = ""
		btn.Pay = false
	}
}

// WithSwitchInlineQuery sets SwitchInlineQuery and defaults everything other.
func WithSwitchInlineQuery(query string) func(*Button) {
	return func(btn *Button) {
		btn.SwitchInlineQuery = query
		btn.URL = ""
		btn.CallbackData = ""
		btn.SwitchInlineQueryCurrentChat = ""
		btn.CallbackGame = ""
		btn.Pay = false
	}
}

// WithSwitchInlineQueryCurrentChat sets SwitchInlineQueryCurrentChat and defaults everything other.
func WithSwitchInlineQueryCurrentChat(query string) func(*Button) {
	return func(btn *Button) {
		btn.SwitchInlineQueryCurrentChat = query
		btn.URL = ""
		btn.CallbackData = ""
		btn.SwitchInlineQuery = ""
		btn.CallbackGame = ""
		btn.Pay = false
	}
}

// WithCallbackGame sets CallbackGame and defaults everything other.
func WithCallbackGame(cbgame string) func(*Button) {
	return func(btn *Button) {
		btn.CallbackGame = cbgame
	}
}

// Pay sets Pay and defaults everything other.
func Pay() func(*Button) {
	return func(btn *Button) {
		btn.URL = ""
		btn.CallbackData = ""
		btn.SwitchInlineQuery = ""
		btn.SwitchInlineQueryCurrentChat = ""
		btn.CallbackGame = ""
		btn.Pay = true
	}
}
