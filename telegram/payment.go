package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SendPaidMedia performs sendPaidMedia telegram api method.
//
// options:
//
// option.SendPaidMediaCaption(), option.SendPaidMediaParseMode(),
// option.SendPaidMediaShowCaptionAboveMedia(), option.SendPaidMediaReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#sendpaidmedia
func (t *Client) SendPaidMedia(ctx context.Context, chatID int64, starCount int, media []method.InputPaidMedia, opts ...method.Option[*method.SendPaidMedia]) (types.Message, error) {
	m := method.NewSendPaidMedia(chatID, starCount, media, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// AnswerShippingQuery performs answerShippingQuery telegram api method.
//
// options:
//
// option.AnswerShippingQueryShippingOptions(), option.AnswerShippingQueryErrorMessage()
//
// Reference: https://core.telegram.org/bots/api#answershippingquery
func (t *Client) AnswerShippingQuery(ctx context.Context, shippingQueryID string, ok bool, opts ...method.Option[*method.AnswerShippingQuery]) (bool, error) {
	m := method.NewAnswerShippingQuery(shippingQueryID, ok, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// AnswerPreCheckoutQuery performs answerPreCheckoutQuery telegram api method.
//
// options:
//
// option.AnswerPreCheckoutQueryErrorMessage()
//
// Reference: https://core.telegram.org/bots/api#answerprecheckoutquery
func (t *Client) AnswerPreCheckoutQuery(ctx context.Context, preCheckoutQueryID string, ok bool, opts ...method.Option[*method.AnswerPreCheckoutQuery]) (bool, error) {
	m := method.NewAnswerPreCheckoutQuery(preCheckoutQueryID, ok, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetStarTransactions performs getStarTransactions telegram api method.
//
// options:
//
// option.GetStarTransactionsOffset(), option.GetStarTransactionsLimit()
//
// Reference: https://core.telegram.org/bots/api#getstartransactions
func (t *Client) GetStarTransactions(ctx context.Context, opts ...method.Option[*method.GetStarTransactions]) (types.StarTransactions, error) {
	m := method.NewGetStarTransactions(opts...)
	var result types.StarTransactions
	err := t.process(ctx, m, &result)
	return result, err
}

// RefundStarPayment performs refundStarPayment telegram api method.
//
// Reference: https://core.telegram.org/bots/api#refundstarpayment
func (t *Client) RefundStarPayment(ctx context.Context, userID int64, telegramPaymentCharge string) (bool, error) {
	m := method.NewRefundStarPayment(userID, telegramPaymentCharge)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
