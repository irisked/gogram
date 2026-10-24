package gogram

import (
	"context"
	"errors"
	"strings"

	"github.com/irisked/gogram/telegram"
)

// Errors returned by Update shortcuts used on an update that lacks
// what they need.
var (
	ErrNoChat             = errors.New("gogram: update has no chat")
	ErrNoCallbackQuery    = errors.New("gogram: update is not a callback query")
	ErrNoInlineQuery      = errors.New("gogram: update is not an inline query")
	ErrNoShippingQuery    = errors.New("gogram: update is not a shipping query")
	ErrNoPreCheckoutQuery = errors.New("gogram: update is not a pre-checkout query")
	ErrNoMessage          = errors.New("gogram: callback query message is unavailable")
)

// Update is the incoming update passed to handlers. It wraps the raw
// telegram.Update, gives access to the Telegram client, and offers
// shortcuts that act on the update's chat or callback query.
//
// The shortcuts take the handler's context.Context as their first
// argument, like Client methods do. Those named after a Client method
// call it for the update's chat; the Reply shortcuts call the send
// methods (Reply calls SendMessage, ReplyWithPhoto calls SendPhoto, and
// so on).
type Update struct {
	client *telegram.Client
	update telegram.Update
}

// NewUpdate wraps update for a handler. Handlers receive one from the
// Bot; NewUpdate is for testing handlers in isolation.
func NewUpdate(client *telegram.Client, update telegram.Update) *Update {
	return &Update{client: client, update: update}
}

// Client returns the Telegram Bot API client.
func (u *Update) Client() *telegram.Client { return u.client }

// Raw returns the raw Bot API update.
func (u *Update) Raw() *telegram.Update { return &u.update }

// Message returns the message the update carries (see
// telegram.Update.AnyMessage), or nil.
func (u *Update) Message() *telegram.Message { return u.update.AnyMessage() }

// CallbackQuery returns the callback query, or nil.
func (u *Update) CallbackQuery() *telegram.CallbackQuery { return u.update.CallbackQuery }

// InlineQuery returns the inline query, or nil.
func (u *Update) InlineQuery() *telegram.InlineQuery { return u.update.InlineQuery }

// Chat returns the chat the update belongs to, or nil (e.g. for inline
// queries).
func (u *Update) Chat() *telegram.Chat {
	r := &u.update
	switch {
	case r.AnyMessage() != nil:
		return r.AnyMessage().Chat
	case r.CallbackQuery != nil:
		if r.CallbackQuery.Message != nil {
			return r.CallbackQuery.Message.Chat
		}
	case r.MyChatMember != nil:
		return r.MyChatMember.Chat
	case r.ChatMember != nil:
		return r.ChatMember.Chat
	case r.ChatJoinRequest != nil:
		return r.ChatJoinRequest.Chat
	case r.MessageReaction != nil:
		return r.MessageReaction.Chat
	case r.MessageReactionCount != nil:
		return r.MessageReactionCount.Chat
	case r.ChatBoost != nil:
		return r.ChatBoost.Chat
	case r.RemovedChatBoost != nil:
		return r.RemovedChatBoost.Chat
	case r.DeletedBusinessMessages != nil:
		return r.DeletedBusinessMessages.Chat
	case r.StoppedMessageGeneration != nil:
		return r.StoppedMessageGeneration.Chat
	}
	return nil
}

// ChatID returns the ID of Chat, or 0 if there is none.
func (u *Update) ChatID() int64 {
	if chat := u.Chat(); chat != nil {
		return chat.ID
	}
	return 0
}

// From returns the user who caused the update, or nil (e.g. for
// channel posts).
func (u *Update) From() *telegram.User {
	r := &u.update
	switch {
	case r.AnyMessage() != nil:
		return r.AnyMessage().From
	case r.CallbackQuery != nil:
		return r.CallbackQuery.From
	case r.InlineQuery != nil:
		return r.InlineQuery.From
	case r.ChosenInlineResult != nil:
		return r.ChosenInlineResult.From
	case r.ShippingQuery != nil:
		return r.ShippingQuery.From
	case r.PreCheckoutQuery != nil:
		return r.PreCheckoutQuery.From
	case r.PurchasedPaidMedia != nil:
		return r.PurchasedPaidMedia.From
	case r.PollAnswer != nil:
		return r.PollAnswer.User
	case r.MyChatMember != nil:
		return r.MyChatMember.From
	case r.ChatMember != nil:
		return r.ChatMember.From
	case r.ChatJoinRequest != nil:
		return r.ChatJoinRequest.From
	case r.MessageReaction != nil:
		return r.MessageReaction.User
	case r.BusinessConnection != nil:
		return r.BusinessConnection.User
	case r.ManagedBot != nil:
		return r.ManagedBot.User
	case r.Subscription != nil:
		return r.Subscription.User
	default:
		return nil
	}
}

// Args returns the arguments of a command ("/ban 42 spam" yields
// ["42", "spam"]) or of a callback query (see telegram.CallbackButton).
// It returns nil for other updates.
func (u *Update) Args() []string {
	if cq := u.update.CallbackQuery; cq != nil {
		return cq.Args()
	}
	if m := u.Message(); m != nil && m.IsCommand() {
		return strings.Fields(m.CommandArguments())
	}
	return nil
}

func (u *Update) chatID() (int64, error) {
	if id := u.ChatID(); id != 0 {
		return id, nil
	}
	return 0, ErrNoChat
}

func (u *Update) callbackTarget() (telegram.MessageTarget, error) {
	cq := u.update.CallbackQuery
	if cq == nil {
		return telegram.MessageTarget{}, ErrNoCallbackQuery
	}
	target, ok := cq.Target()
	if !ok {
		return telegram.MessageTarget{}, ErrNoMessage
	}
	return target, nil
}

// --- Answering the update's query ---

// AnswerCallbackQuery answers the update's callback query.
func (u *Update) AnswerCallbackQuery(ctx context.Context, opts ...telegram.AnswerCallbackQueryOption) error {
	q := u.update.CallbackQuery
	if q == nil {
		return ErrNoCallbackQuery
	}
	return u.client.AnswerCallbackQuery(ctx, q.ID, opts...)
}

// AnswerInlineQuery answers the update's inline query.
func (u *Update) AnswerInlineQuery(ctx context.Context, results []telegram.InlineQueryResult, opts ...telegram.AnswerInlineQueryOption) error {
	q := u.update.InlineQuery
	if q == nil {
		return ErrNoInlineQuery
	}
	return u.client.AnswerInlineQuery(ctx, q.ID, results, opts...)
}

// AnswerShippingQuery answers the update's shipping query.
func (u *Update) AnswerShippingQuery(ctx context.Context, ok bool, opts ...telegram.AnswerShippingQueryOption) error {
	q := u.update.ShippingQuery
	if q == nil {
		return ErrNoShippingQuery
	}
	return u.client.AnswerShippingQuery(ctx, q.ID, ok, opts...)
}

// AnswerPreCheckoutQuery answers the update's pre-checkout query.
func (u *Update) AnswerPreCheckoutQuery(ctx context.Context, ok bool, opts ...telegram.AnswerPreCheckoutQueryOption) error {
	q := u.update.PreCheckoutQuery
	if q == nil {
		return ErrNoPreCheckoutQuery
	}
	return u.client.AnswerPreCheckoutQuery(ctx, q.ID, ok, opts...)
}
