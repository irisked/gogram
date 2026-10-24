package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
)

// AnswerInlineQuery performs answerInlineQuery telegram api method.
//
// options:
//
// option.AnswerInlineQueryCacheTime(), option.AnswerInlineQueryIsPersonal(),
// option.AnswerInlineQueryNextOffset(), option.AnswerInlineQueryButton()
//
// Reference: https://core.telegram.org/bots/api#answerinlinequery
func (t *Client) AnswerInlineQuery(ctx context.Context, inlineQueryID string, results []method.InlineQueryResult, opts ...method.Option[*method.AnswerInlineQuery]) (bool, error) {
	m := method.NewAnswerInlineQuery(inlineQueryID, results, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// AnswerWebAppQuery performs answerWebAppQuery telegram api method.
//
// Reference: https://core.telegram.org/bots/api#answerwebappquery
func (t *Client) AnswerWebAppQuery(ctx context.Context, webAppQueryID, result string) (method.SentWebAppMessage, error) {
	m := method.NewAnswerWebAppQuery(webAppQueryID, result)
	var resultSent method.SentWebAppMessage
	err := t.process(ctx, m, &resultSent)
	return resultSent, err
}
