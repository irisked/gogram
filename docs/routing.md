# Routing

A `*gogram.Bot` sends each incoming update to one handler. You register handlers with the `On*` methods; the bot picks the most specific one that matches.

## Routes

Every update has a route, built from up to three parts:

| Update | Route |
|---|---|
| Text message | `message` → `text` |
| Command `/start ref42` | `message` → `command` → `start` |
| Photo message | `message` → `photo` |
| Edited message | `edited_message` → `text` (or another message type) |
| Button press with data `vote:up` | `callback_query` → `vote` |
| Inline query | `inline_query` |
| Any other update type | its type, e.g. `poll_answer`, `my_chat_member` |

The bot looks for a handler for the full route first, then drops the last part and tries again, and finally falls back to `OnAny`. Parts are matched whole: a handler for the command `start` doesn't receive `/startgame`.

Example: with these handlers registered

```go
bot.OnStart(start)     // message → command → start
bot.OnText(text)       // message → text
bot.OnMessage(message) // message
bot.OnAny(any)
```

| Incoming | Handled by |
|---|---|
| `/start` | `start` |
| `/startgame` | `message` (no handler for that command) |
| `hello` | `text` |
| a photo | `message` |
| an edited message | `any` |
| an inline query | `any` |

An update with no matching handler, and no `OnAny`, is dropped.

## Messages

| Method | Receives |
|---|---|
| `OnMessage` | every new message without a more specific handler |
| `OnText` | text messages that aren't commands for this bot |
| `OnCommand(name)` | the command `/name` |
| `OnStart`, `OnHelp` | `/start`, `/help` |
| `OnPhoto`, `OnVideo`, `OnAnimation`, `OnAudio`, `OnDocument`, `OnVoice`, `OnVideoNote`, `OnSticker` | media messages |
| `OnContact`, `OnLocation`, `OnVenue` | shared contacts and places |
| `OnNewChatMembers`, `OnLeftChatMember`, `OnNewChatPhoto`, `OnDeleteChatPhoto` | service messages |
| `OnInvoice`, `OnSuccessfulPayment` | payment messages |
| `OnMessageType(t)` | any `telegram.MessageType`, e.g. `telegram.MessageTypePoll` |

These all apply to new messages only. Edited messages, channel posts and business messages have their own update types: `OnEditedMessage`, `OnChannelPost`, `OnEditedChannelPost`, `OnBusinessMessage`.

A message has exactly one type, given by `telegram.Message.Type()`. A message that starts with a command is a command, not text. A venue is a venue, not a location, and an animation is an animation, not a document.

### Commands

```go
bot.OnCommand("settings", func(ctx context.Context, u *gogram.Update) error {
	args := u.Args() // "/settings lang en" → ["lang", "en"]
	// ...
	return nil
})
```

- The leading slash is optional when registering: `OnCommand("/settings", h)` is the same.
- Commands are matched case-insensitively.
- In groups, commands can name a bot: `/settings@my_bot`. Commands naming this bot are handled as commands; commands naming another bot are routed as text, so two bots in a group don't both answer.

To tell those apart the bot needs its own username. `StartPolling` and `WebhookHandler` look it up with `getMe`; set it yourself with `gogram.WithUsername("my_bot")` to skip the lookup. While it is unknown, for example with `HandleUpdate` called directly, commands naming any bot are accepted.

## Callback queries

Callback queries are button presses on inline keyboards. Buttons made with `telegram.CallbackButton(text, route, args...)` carry `route:arg1,arg2` as their data, and `OnCallback(route)` receives them:

```go
kb := telegram.NewInlineKeyboard(
	telegram.CallbackButton("Delete", "delete", "42"),
)

bot.OnCallback("delete", func(ctx context.Context, u *gogram.Update) error {
	id := u.Args()[0] // "42"
	// ...
	return u.AnswerCallbackQuery(ctx)
})
```

`OnCallbackQuery` receives presses whose route has no handler of its own. See [Keyboards and callbacks](keyboards.md).

## Other updates

| Method | Receives |
|---|---|
| `OnInlineQuery` | inline queries (`@yourbot query` typed in any chat) |
| `OnChosenInlineResult` | inline results the user picked |
| `OnPoll`, `OnPollAnswer` | poll state changes, answers to non-anonymous polls |
| `OnMyChatMember` | the bot was added to, removed from or blocked in a chat |
| `OnChatMember` | other members' status changed (must be requested) |
| `OnChatJoinRequest` | requests to join a chat the bot administers |
| `OnMessageReaction` | reactions to messages (must be requested) |
| `OnShippingQuery`, `OnPreCheckoutQuery` | payment flow queries |
| `OnUpdate(t)` | all updates of a `telegram.UpdateType`, e.g. `telegram.UpdateChatBoost` |
| `OnAny` | everything without a more specific handler |

Telegram doesn't send every update type by default: `chat_member`, `message_reaction` and `message_reaction_count` must be requested with `gogram.WithAllowedUpdates` (see [Polling and webhooks](updates.md)).

## Registration rules

- Register all handlers before calling `StartPolling` or serving `WebhookHandler`. Registering while updates are being handled is not safe.
- Registering a second handler for the same route replaces the first.
- Updates from different chats are handled concurrently, so handlers that share state must synchronize access to it. Updates from the same chat are handled one at a time, in order.
- A handler that panics is recovered, and the panic goes to the error handler (see [Errors](errors.md)).
