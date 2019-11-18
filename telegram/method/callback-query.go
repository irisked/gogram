package method

// AnswerCallbackQuery contains information about answerCallbackQuery telegram api method.
type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
	URL             string `json:"url,omitempty"`
	SimpleMethod
}

// AnswerCallbackQueryOption interface for setting AnswerCallbackQuery optional parameters.
type AnswerCallbackQueryOption interface {
	ApplyAnswerCallbackQueryOption(*AnswerCallbackQuery)
}

// NewAnswerCallbackQuery creates AnswerCallbackQuery.
func NewAnswerCallbackQuery(id string, options ...AnswerCallbackQueryOption) *AnswerCallbackQuery {
	method := new(AnswerCallbackQuery)
	method.CallbackQueryID = id
	for _, option := range options {
		option.ApplyAnswerCallbackQueryOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *AnswerCallbackQuery) Method() string {
	return "answerCallbackQuery"
}
