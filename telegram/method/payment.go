package method

import ()

// InputPaidMedia contains information about input paid media.
type InputPaidMedia struct {
	Type  string `json:"type" form:"type"`
	Photo string `json:"photo,omitempty" form:"photo,omitempty"`
	Video string `json:"video,omitempty" form:"video,omitempty"`
}

// SendPaidMedia contains information about sendPaidMedia telegram api method.
type SendPaidMedia struct {
	ChatID                 int64            `json:"chat_id" form:"chat_id"`
	StarCount              int              `json:"star_count" form:"star_count"`
	Media                  []InputPaidMedia `json:"media" form:"media"`
	Caption                string           `json:"caption,omitempty" form:"caption,omitempty"`
	ParseMode              string           `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	ShowCapitionAboveMedia *bool            `json:"show_caption_above_media,omitempty" form:"show_caption_above_media,omitempty"`
	ReplyMarkup            string           `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	SimpleMethod
}

// NewSendPaidMedia creates SendPaidMedia.
func NewSendPaidMedia(chatID int64, starCount int, media []InputPaidMedia, opts ...Option[*SendPaidMedia]) *SendPaidMedia {
	m := &SendPaidMedia{
		ChatID:    chatID,
		StarCount: starCount,
		Media:     media,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SendPaidMedia) Method() string {
	return "sendPaidMedia"
}

// SetCaption sets caption
func (r *SendPaidMedia) SetCaption(c string) {
	r.Caption = c
}

// SetParseMode sets parse_mode
func (r *SendPaidMedia) SetParseMode(p string) {
	r.ParseMode = p
}

// SetShowCaptionAboveMedia sets show_caption_above_media
func (r *SendPaidMedia) SetShowCaptionAboveMedia(s *bool) {
	r.ShowCapitionAboveMedia = s
}

// SetReplyMarkup sets reply_markup
func (r *SendPaidMedia) SetReplyMarkup(r2 string) {
	r.ReplyMarkup = r2
}

// AnswerShippingQuery contains information about answerShippingQuery telegram api method.
type AnswerShippingQuery struct {
	ShippingQueryID string `json:"shipping_query_id" form:"shipping_query_id"`
	OK              bool   `json:"ok" form:"ok"`
	ShippingOptions string `json:"shipping_options,omitempty" form:"shipping_options,omitempty"`
	ErrorMessag     string `json:"error_message,omitempty" form:"error_message,omitempty"`
	SimpleMethod
}

// NewAnswerShippingQuery creates AnswerShippingQuery.
func NewAnswerShippingQuery(shippingQueryID string, ok bool, opts ...Option[*AnswerShippingQuery]) *AnswerShippingQuery {
	m := &AnswerShippingQuery{
		ShippingQueryID: shippingQueryID,
		OK:              ok,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *AnswerShippingQuery) Method() string {
	return "answerShippingQuery"
}

// SetShippingOptions sets shipping_options
func (r *AnswerShippingQuery) SetShippingOptions(s string) {
	r.ShippingOptions = s
}

// SetErrorMessage sets error_message
func (r *AnswerShippingQuery) SetErrorMessage(e string) {
	r.ErrorMessag = e
}

// AnswerPreCheckoutQuery contains information about answerPreCheckoutQuery telegram api method.
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id" form:"pre_checkout_query_id"`
	OK                 bool   `json:"ok" form:"ok"`
	ErrorMessag        string `json:"error_message,omitempty" form:"error_message,omitempty"`
	SimpleMethod
}

// NewAnswerPreCheckoutQuery creates AnswerPreCheckoutQuery.
func NewAnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, opts ...Option[*AnswerPreCheckoutQuery]) *AnswerPreCheckoutQuery {
	m := &AnswerPreCheckoutQuery{
		PreCheckoutQueryID: preCheckoutQueryID,
		OK:                 ok,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *AnswerPreCheckoutQuery) Method() string {
	return "answerPreCheckoutQuery"
}

// SetErrorMessage sets error_message
func (r *AnswerPreCheckoutQuery) SetErrorMessage(e string) {
	r.ErrorMessag = e
}

// GetStarTransactions contains information about getStarTransactions telegram api method.
type GetStarTransactions struct {
	Offset int `json:"offset,omitempty" form:"offset,omitempty"`
	Limit  int `json:"limit,omitempty" form:"limit,omitempty"`
	SimpleMethod
}

// NewGetStarTransactions creates GetStarTransactions.
func NewGetStarTransactions(opts ...Option[*GetStarTransactions]) *GetStarTransactions {
	m := &GetStarTransactions{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetStarTransactions) Method() string {
	return "getStarTransactions"
}

// SetOffset sets offset
func (r *GetStarTransactions) SetOffset(o int) {
	r.Offset = o
}

// SetLimit sets limit
func (r *GetStarTransactions) SetLimit(l int) {
	r.Limit = l
}

// RefundStarPayment contains information about refundStarPayment telegram api method.
type RefundStarPayment struct {
	UserID                int64  `json:"user_id" form:"user_id"`
	TelegramPaymentCharge string `json:"telegram_payment_charge" form:"telegram_payment_charge"`
	SimpleMethod
}

// NewRefundStarPayment creates RefundStarPayment.
func NewRefundStarPayment(userID int64, telegramPaymentCharge string) *RefundStarPayment {
	m := &RefundStarPayment{
		UserID:                userID,
		TelegramPaymentCharge: telegramPaymentCharge,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *RefundStarPayment) Method() string {
	return "refundStarPayment"
}

// SentWebAppMessage contains information about an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
