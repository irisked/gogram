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

// NewAnswerCallbackQuery creates AnswerCallbackQuery.
func NewAnswerCallbackQuery(id string, opts ...Option[*AnswerCallbackQuery]) *AnswerCallbackQuery {
	method := new(AnswerCallbackQuery)
	method.CallbackQueryID = id
	for _, o := range opts {
		o(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *AnswerCallbackQuery) Method() string {
	return "answerCallbackQuery"
}

// SetText sets text
func (r *AnswerCallbackQuery) SetText(s string) {
	r.Text = s
}

// SetShowAlert sets show_alert
func (r *AnswerCallbackQuery) SetShowAlert(v bool) {
	r.ShowAlert = v
}

// SetCacheTime sets cache_time
func (r *AnswerCallbackQuery) SetCacheTime(v int) {
	r.CacheTime = v
}

// SetURL sets url
func (r *AnswerCallbackQuery) SetURL(s string) {
	r.URL = s
}
