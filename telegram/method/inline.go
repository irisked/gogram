package method

import ()

// InlineQueryResult contains information about inline query result.
type InlineQueryResult struct {
	Type                string `json:"type" form:"type"`
	ID                  string `json:"id" form:"id"`
	Title               string `json:"title,omitempty" form:"title,omitempty"`
	Caption             string `json:"caption,omitempty" form:"caption,omitempty"`
	InputMessageContent string `json:"input_message_content,omitempty" form:"input_message_content,omitempty"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	URL                 string `json:"url,omitempty" form:"url,omitempty"`
	ThumbURL            string `json:"thumb_url,omitempty" form:"thumb_url,omitempty"`
	ThumbWidth          int    `json:"thumb_width,omitempty" form:"thumb_width,omitempty"`
	ThumbHeight         int    `json:"thumb_height,omitempty" form:"thumb_height,omitempty"`
}

// AnswerInlineQuery contains information about answerInlineQuery telegram api method.
type AnswerInlineQuery struct {
	InlineQueryID string              `json:"inline_query_id" form:"inline_query_id"`
	Results       []InlineQueryResult `json:"results" form:"results"`
	CacheTime     int                 `json:"cache_time,omitempty" form:"cache_time,omitempty"`
	IsPersonal    bool                `json:"is_personal,omitempty" form:"is_personal,omitempty"`
	NextOffset    string              `json:"next_offset,omitempty" form:"next_offset,omitempty"`
	Button        string              `json:"button,omitempty" form:"button,omitempty"`
	SimpleMethod
}

// NewAnswerInlineQuery creates AnswerInlineQuery.
func NewAnswerInlineQuery(inlineQueryID string, results []InlineQueryResult, opts ...Option[*AnswerInlineQuery]) *AnswerInlineQuery {
	m := &AnswerInlineQuery{
		InlineQueryID: inlineQueryID,
		Results:       results,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *AnswerInlineQuery) Method() string {
	return "answerInlineQuery"
}

// SetCacheTime sets cache_time
func (r *AnswerInlineQuery) SetCacheTime(c int) {
	r.CacheTime = c
}

// SetIsPersonal sets is_personal
func (r *AnswerInlineQuery) SetIsPersonal(i bool) {
	r.IsPersonal = i
}

// SetNextOffset sets next_offset
func (r *AnswerInlineQuery) SetNextOffset(n string) {
	r.NextOffset = n
}

// SetButton sets button
func (r *AnswerInlineQuery) SetButton(b string) {
	r.Button = b
}

// AnswerWebAppQuery contains information about answerWebAppQuery telegram api method.
type AnswerWebAppQuery struct {
	WebAppQueryID string `json:"web_app_query_id" form:"web_app_query_id"`
	Result        string `json:"result" form:"result"`
	SimpleMethod
}

// NewAnswerWebAppQuery creates AnswerWebAppQuery.
func NewAnswerWebAppQuery(webAppQueryID, result string) *AnswerWebAppQuery {
	m := &AnswerWebAppQuery{
		WebAppQueryID: webAppQueryID,
		Result:        result,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *AnswerWebAppQuery) Method() string {
	return "answerWebAppQuery"
}

// SetResult sets result
func (r *AnswerWebAppQuery) SetResult(r2 string) {
	r.Result = r2
}
