# The update

Every handler receives a `context.Context` and a `*gogram.Update`, and returns an error:

```go
type Handler func(ctx context.Context, u *gogram.Update) error
```

The update wraps the raw `telegram.Update`, gives access to the Telegram client, and has shortcuts that act on the update's chat or callback query. The error goes to the bot's error handler (see [Errors](errors.md)).

## The context

`ctx` carries the deadline, cancellation and values of the update's handling. Pass it to every shortcut and `telegram.Client` call the handler makes, and to anything else that takes a context, such as database queries:

```go
bot.OnCommand("pin", func(ctx context.Context, u *gogram.Update) error {
	reply := u.Message().ReplyToMessage
	if reply == nil {
		return nil
	}
	return u.PinChatMessage(ctx, reply.MessageID)
})
```

For webhooks, `ctx` ends when Telegram's HTTP request does. During polling shutdown, running handlers keep a context that is *not* cancelled, so they can finish (see [Polling and webhooks](updates.md)).

Middleware passes values to handlers through the context; see [Middleware](middleware.md#passing-values-to-handlers).

## Reading the update

| Method | Returns |
|---|---|
| `Raw()` | the raw `*telegram.Update`, with every kind of update |
| `Message()` | the message of a message, edited message, channel post, business message or guest message update; otherwise nil |
| `CallbackQuery()`, `InlineQuery()` | that part of the update, or nil |
| `Chat()` | the chat the update happened in, or nil (inline queries have none) |
| `ChatID()` | `Chat().ID`, or 0 |
| `From()` | the user who caused the update, or nil (channel posts have none) |
| `Args()` | command arguments (`/ban 42 spam` → `["42", "spam"]`) or callback arguments (`vote:up` → `["up"]`) |
| `Client()` | the `*telegram.Client` |

For a button press, `From()` is the user who pressed it, and `Chat()` is the chat of the message the button is attached to. `Chat()` and `From()` also cover chat member updates, join requests, reactions, poll answers, boosts and the other update kinds.

## Shortcuts

Shortcuts call a `telegram.Client` method with the update's chat or message filled in. Each takes `ctx` first, then the method's other required parameters and options, and returns what the method returns.

### Sending to the update's chat

Every send method has a `Reply` shortcut: `Reply` calls `SendMessage`, and `ReplyWith<X>` calls `Send<X>`.

```go
u.Reply(ctx, text, opts...)
u.ReplyWithPhoto(ctx, file, opts...)
u.ReplyWithDocument(ctx, file, opts...)
u.ReplyWithMediaGroup(ctx, media, opts...)
u.ReplyWithLocation(ctx, latitude, longitude, opts...)
u.ReplyWithPoll(ctx, question, options, opts...)
u.ReplyWithDice(ctx, opts...)
u.ReplyWithInvoice(ctx, title, description, payload, currency, prices, opts...)
// ... and ReplyWithAudio, ReplyWithVideo, ReplyWithAnimation, ReplyWithVoice,
// ReplyWithVideoNote, ReplyWithSticker, ReplyWithVenue, ReplyWithContact, ...
```

`Reply` sends a new message to the chat; it doesn't quote the incoming one. To quote it:

```go
u.Reply(ctx, telegram.Text("this one?"), telegram.WithReplyToMessageID(u.Message().MessageID))
```

### Acting on the update's chat

```go
u.SendChatAction(ctx, telegram.ChatActionTyping)
u.DeleteMessage(ctx, messageID)
u.DeleteMessages(ctx, messageIDs)
u.SetMessageReaction(ctx, messageID, telegram.WithReaction(&telegram.ReactionTypeEmoji{Emoji: "👍"}))
u.BanChatMember(ctx, userID, opts...)
u.UnbanChatMember(ctx, userID, opts...)
u.RestrictChatMember(ctx, userID, permissions, opts...)
u.PromoteChatMember(ctx, userID, opts...)
u.ApproveChatJoinRequest(ctx, userID)
u.DeclineChatJoinRequest(ctx, userID)
u.PinChatMessage(ctx, messageID, opts...)
u.UnpinChatMessage(ctx, opts...)
u.UnpinAllChatMessages(ctx)
u.GetChat(ctx)
u.GetChatMember(ctx, userID)
u.GetChatAdministrators(ctx)
u.LeaveChat(ctx)
```

### Answering the update's query

```go
u.AnswerCallbackQuery(ctx, opts...)         // a button press
u.AnswerInlineQuery(ctx, results, opts...)  // an inline query
u.AnswerShippingQuery(ctx, ok, opts...)     // a shipping query
u.AnswerPreCheckoutQuery(ctx, ok, opts...)  // a pre-checkout query
```

### Editing the message a button belongs to

In a callback query handler, these edit the message that has the pressed button. That can be a chat message or a message sent in inline mode:

```go
u.EditMessageText(ctx, text, opts...)
u.EditMessageCaption(ctx, telegram.WithCaption(caption))
u.EditMessageMedia(ctx, media, opts...)
u.EditMessageReplyMarkup(ctx, telegram.WithInlineKeyboard(kb)) // no keyboard removes it
u.EditMessageLiveLocation(ctx, latitude, longitude, opts...)
u.StopMessageLiveLocation(ctx, opts...)
```

They return the edited `*telegram.Message`, or nil for inline messages; Telegram doesn't return those.

Anything else is available through `u.Client()`.

## When a shortcut doesn't apply

Shortcuts return an error instead of panicking when the update lacks what they need:

| Error | When |
|---|---|
| `gogram.ErrNoChat` | a chat shortcut is used on an update without a chat, e.g. an inline query |
| `gogram.ErrNoCallbackQuery` | a callback shortcut is used on another update type |
| `gogram.ErrNoInlineQuery`, `ErrNoShippingQuery`, `ErrNoPreCheckoutQuery` | an answer shortcut is used on another update type |
| `gogram.ErrNoMessage` | an edit shortcut is used on a button press whose message Telegram no longer delivers |

```go
if _, err := u.EditMessageText(ctx, telegram.Text("done")); errors.Is(err, gogram.ErrNoMessage) {
	_, err = u.Reply(ctx, telegram.Text("done"))
	return err
}
```
