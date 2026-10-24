package telegram

import (
	"context"

	"github.com/irisked/gogram/telegram/method"
)

// AnswerCallbackQuery performs answerCallbackQuery telegram api method.
//
// options:
//
// option.Text(text string), option.ShowAlert(), option.URL(url string),
// option.CacheTime(time int)
//
// Reference: https://core.telegram.org/bots/api#answercallbackquery
func (t *Client) AnswerCallbackQuery(ctx context.Context, id string, opts ...method.Option[*method.AnswerCallbackQuery]) (bool, error) {
	m := method.NewAnswerCallbackQuery(id, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}
