package gogram

import (
	"strings"

	"github.com/irisked/gogram/telegram"
)

// Routes are slash-separated: "/<update type>", then for messages
// "/<message type>" and for commands "/<command>", and for callback
// queries "/<callback route>". Handlers registered for a shorter route
// catch everything below it that has no handler of its own.

func route(segments ...string) string {
	return "/" + strings.Join(segments, "/")
}

// routeOf returns the route of an update. A command addressed to
// another bot ("/start@other_bot") is routed as a text message; while
// the bot's username is unknown, commands to any bot are accepted.
func routeOf(u *telegram.Update, username string) string {
	ut := u.Type()
	if ut == "" {
		return "/"
	}
	switch m := u.AnyMessage(); {
	case m != nil:
		mt := m.Type()
		if mt != telegram.MessageTypeCommand {
			return route(ut.String(), mt.String())
		}
		if bot := m.CommandBot(); bot != "" && username != "" && !strings.EqualFold(bot, username) {
			return route(ut.String(), telegram.MessageTypeText.String())
		}
		return route(ut.String(), mt.String(), strings.ToLower(m.Command()))
	case u.CallbackQuery != nil:
		return route(ut.String(), u.CallbackQuery.Route())
	default:
		return route(ut.String())
	}
}

func (b *Bot) register(path string, h Handler) {
	b.routes.Insert(path, h)
}

// OnAny handles every update that no more specific handler matches.
func (b *Bot) OnAny(h Handler) { b.register("", h) }

// OnUpdate handles all updates of the given type.
func (b *Bot) OnUpdate(ut telegram.UpdateType, h Handler) { b.register(route(ut.String()), h) }

// OnMessage handles new messages.
func (b *Bot) OnMessage(h Handler) { b.OnUpdate(telegram.UpdateMessage, h) }

// OnEditedMessage handles edited messages.
func (b *Bot) OnEditedMessage(h Handler) { b.OnUpdate(telegram.UpdateEditedMessage, h) }

// OnChannelPost handles new channel posts.
func (b *Bot) OnChannelPost(h Handler) { b.OnUpdate(telegram.UpdateChannelPost, h) }

// OnEditedChannelPost handles edited channel posts.
func (b *Bot) OnEditedChannelPost(h Handler) { b.OnUpdate(telegram.UpdateEditedChannelPost, h) }

// OnInlineQuery handles inline queries.
func (b *Bot) OnInlineQuery(h Handler) { b.OnUpdate(telegram.UpdateInlineQuery, h) }

// OnChosenInlineResult handles inline results chosen by users.
func (b *Bot) OnChosenInlineResult(h Handler) { b.OnUpdate(telegram.UpdateChosenInlineResult, h) }

// OnCallbackQuery handles all callback queries.
func (b *Bot) OnCallbackQuery(h Handler) { b.OnUpdate(telegram.UpdateCallbackQuery, h) }

// OnCallback handles callback queries from buttons created with
// telegram.CallbackButton(text, route, ...). Read the button arguments
// with Update.Args.
func (b *Bot) OnCallback(name string, h Handler) {
	b.register(route(telegram.UpdateCallbackQuery.String(), name), h)
}

// OnPoll handles updates about the state of polls the bot sent.
func (b *Bot) OnPoll(h Handler) { b.OnUpdate(telegram.UpdatePoll, h) }

// OnPollAnswer handles answers to non-anonymous polls the bot sent.
func (b *Bot) OnPollAnswer(h Handler) { b.OnUpdate(telegram.UpdatePollAnswer, h) }

// OnMyChatMember handles changes of the bot's own membership in a chat,
// such as being added to a group or blocked by a user.
func (b *Bot) OnMyChatMember(h Handler) { b.OnUpdate(telegram.UpdateMyChatMember, h) }

// OnChatMember handles changes of other members' status. Telegram sends
// them only to administrators, and only if requested with
// WithAllowedUpdates.
func (b *Bot) OnChatMember(h Handler) { b.OnUpdate(telegram.UpdateChatMember, h) }

// OnChatJoinRequest handles requests to join a chat the bot administers.
func (b *Bot) OnChatJoinRequest(h Handler) { b.OnUpdate(telegram.UpdateChatJoinRequest, h) }

// OnMessageReaction handles reaction changes on messages. Telegram sends
// them only if requested with WithAllowedUpdates.
func (b *Bot) OnMessageReaction(h Handler) { b.OnUpdate(telegram.UpdateMessageReaction, h) }

// OnBusinessMessage handles messages in a connected business account.
func (b *Bot) OnBusinessMessage(h Handler) { b.OnUpdate(telegram.UpdateBusinessMessage, h) }

// OnShippingQuery handles shipping queries.
func (b *Bot) OnShippingQuery(h Handler) { b.OnUpdate(telegram.UpdateShippingQuery, h) }

// OnPreCheckoutQuery handles pre-checkout queries.
func (b *Bot) OnPreCheckoutQuery(h Handler) { b.OnUpdate(telegram.UpdatePreCheckoutQuery, h) }

// OnMessageType handles new messages of the given type.
func (b *Bot) OnMessageType(mt telegram.MessageType, h Handler) {
	b.register(route(telegram.UpdateMessage.String(), mt.String()), h)
}

// OnCommand handles a bot command such as "/settings", given with or
// without the slash. Commands are matched case-insensitively, with or
// without a "@botname" suffix naming this bot; commands for other bots
// are routed as text. Read the arguments with Update.Args.
func (b *Bot) OnCommand(command string, h Handler) {
	command = strings.ToLower(strings.TrimPrefix(command, "/"))
	b.register(route(telegram.UpdateMessage.String(), telegram.MessageTypeCommand.String(), command), h)
}

// OnStart handles the /start command.
func (b *Bot) OnStart(h Handler) { b.OnCommand("start", h) }

// OnHelp handles the /help command.
func (b *Bot) OnHelp(h Handler) { b.OnCommand("help", h) }

// OnText handles text messages that are not commands for this bot.
func (b *Bot) OnText(h Handler) { b.OnMessageType(telegram.MessageTypeText, h) }

// OnAnimation handles animation messages.
func (b *Bot) OnAnimation(h Handler) { b.OnMessageType(telegram.MessageTypeAnimation, h) }

// OnAudio handles audio messages.
func (b *Bot) OnAudio(h Handler) { b.OnMessageType(telegram.MessageTypeAudio, h) }

// OnDocument handles document messages.
func (b *Bot) OnDocument(h Handler) { b.OnMessageType(telegram.MessageTypeDocument, h) }

// OnPhoto handles photo messages.
func (b *Bot) OnPhoto(h Handler) { b.OnMessageType(telegram.MessageTypePhoto, h) }

// OnSticker handles sticker messages.
func (b *Bot) OnSticker(h Handler) { b.OnMessageType(telegram.MessageTypeSticker, h) }

// OnVideo handles video messages.
func (b *Bot) OnVideo(h Handler) { b.OnMessageType(telegram.MessageTypeVideo, h) }

// OnVideoNote handles video note messages.
func (b *Bot) OnVideoNote(h Handler) { b.OnMessageType(telegram.MessageTypeVideoNote, h) }

// OnVoice handles voice messages.
func (b *Bot) OnVoice(h Handler) { b.OnMessageType(telegram.MessageTypeVoice, h) }

// OnContact handles shared contacts.
func (b *Bot) OnContact(h Handler) { b.OnMessageType(telegram.MessageTypeContact, h) }

// OnLocation handles shared locations.
func (b *Bot) OnLocation(h Handler) { b.OnMessageType(telegram.MessageTypeLocation, h) }

// OnVenue handles shared venues.
func (b *Bot) OnVenue(h Handler) { b.OnMessageType(telegram.MessageTypeVenue, h) }

// OnNewChatMembers handles service messages about users joining.
func (b *Bot) OnNewChatMembers(h Handler) { b.OnMessageType(telegram.MessageTypeNewChatMembers, h) }

// OnLeftChatMember handles service messages about a user leaving.
func (b *Bot) OnLeftChatMember(h Handler) { b.OnMessageType(telegram.MessageTypeLeftChatMember, h) }

// OnNewChatPhoto handles service messages about a new chat photo.
func (b *Bot) OnNewChatPhoto(h Handler) { b.OnMessageType(telegram.MessageTypeNewChatPhoto, h) }

// OnDeleteChatPhoto handles service messages about a deleted chat photo.
func (b *Bot) OnDeleteChatPhoto(h Handler) { b.OnMessageType(telegram.MessageTypeDeleteChatPhoto, h) }

// OnInvoice handles invoice messages.
func (b *Bot) OnInvoice(h Handler) { b.OnMessageType(telegram.MessageTypeInvoice, h) }

// OnSuccessfulPayment handles service messages about a successful payment.
func (b *Bot) OnSuccessfulPayment(h Handler) {
	b.OnMessageType(telegram.MessageTypeSuccessfulPayment, h)
}
