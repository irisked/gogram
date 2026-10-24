# Keyboards and callbacks

Telegram has two kinds of keyboards:

- **Inline keyboards** are buttons attached to a message. Pressing one sends your bot a callback query or opens a URL.
- **Reply keyboards** replace the user's keyboard. Pressing a button sends its text as a normal message, or shares a contact or location.

Both are attached with `WithReplyMarkup`.

## Inline keyboards

```go
kb := telegram.NewInlineKeyboard(
	telegram.CallbackButton("👍", "vote", "up"),
	telegram.CallbackButton("👎", "vote", "down"),
).Row(
	telegram.URLButton("Read the rules", "https://example.com/rules"),
)

u.Reply(ctx, telegram.Text("Do you like it?"), telegram.WithReplyMarkup(kb))
```

`NewInlineKeyboard` creates the first row; each `Row` call adds another.

| Button | When pressed |
|---|---|
| `CallbackButton(text, route, args...)` | sends a callback query to the bot |
| `URLButton(text, url)` | opens a URL |
| `WebAppButton(text, url)` | opens a Web App |
| `SwitchInlineQueryButton(text, query)` | lets the user pick a chat and starts an inline query there |
| `SwitchInlineQueryCurrentChatButton(text, query)` | starts an inline query in the current chat |
| `GameButton(text)` | launches the game of a `SendGame` message (must be first) |
| `PayButton(text)` | pays an invoice (must be first) |

You can also build `telegram.InlineKeyboardButton` values yourself, for example to set a button's `Style` or icon.

## Handling button presses

A callback button's data is `route` followed by `:` and its comma-separated arguments, e.g. `vote:up`. `OnCallback(route)` receives presses for that route, and `u.Args()` returns the arguments:

```go
bot.OnCallback("vote", func(ctx context.Context, u *gogram.Update) error {
	choice := u.Args()[0] // "up" or "down"

	// Always answer: the client shows a spinner until you do.
	if err := u.AnswerCallbackQuery(ctx, telegram.WithText("Thanks!")); err != nil {
		return err
	}

	// Replace the question and remove the buttons.
	_, err := u.EditMessageText(ctx, telegram.Text("You voted "+choice))
	return err
})
```

Options for `AnswerCallbackQuery`:

| Option | Effect |
|---|---|
| `WithText(s)` | show a short notification |
| `WithShowAlert()` | show it as an alert with an OK button |
| `WithURL(url)` | open a URL (games, `t.me/yourbot?start=...` links) |
| `WithCacheTime(seconds)` | let the client cache the answer |

Keep in mind:

- Callback data is limited to 64 bytes (`telegram.MaxCallbackData`). A request with a longer one fails before it is sent, with an error naming the button. Store large state on your side and put an ID in the button.
- Routes and arguments can't contain `:` or `,`.
- For buttons made another way, `u.CallbackQuery().Data` holds the raw data.

## Updating a keyboard

The edit methods accept only inline keyboards, so they take `WithInlineKeyboard` rather than `WithReplyMarkup`:

```go
u.EditMessageReplyMarkup(ctx, telegram.WithInlineKeyboard(newKeyboard)) // replace
u.EditMessageReplyMarkup(ctx)                                           // remove

u.EditMessageText(ctx, telegram.Text("Step 2"), telegram.WithInlineKeyboard(step2Keyboard))
```

`WithInlineKeyboard` also works with the send methods.

## Reply keyboards

```go
kb := telegram.NewReplyKeyboard(
	telegram.TextButton("Yes"),
	telegram.TextButton("No"),
).Row(
	telegram.ContactButton("Share my phone"),
	telegram.LocationButton("Share my location"),
)
kb.ResizeKeyboard = true
kb.OneTimeKeyboard = true

u.Reply(ctx, telegram.Text("Choose:"), telegram.WithReplyMarkup(kb))
```

Pressing a text button sends its text, so `OnText` handles it. Shared contacts and locations arrive at `OnContact` and `OnLocation`.

Other fields of `telegram.ReplyKeyboardMarkup`: `IsPersistent`, `InputFieldPlaceholder`, `Selective`.

To remove a reply keyboard, send a message with `telegram.RemoveKeyboard()`:

```go
u.Reply(ctx, telegram.Text("Thanks!"), telegram.WithReplyMarkup(telegram.RemoveKeyboard()))
```

## Forcing a reply

`telegram.NewForceReply()` makes the client open a reply to your message, which is useful for asking a question in a group:

```go
u.Reply(ctx, telegram.Text("What's your name?"), telegram.WithReplyMarkup(telegram.NewForceReply()))
```
