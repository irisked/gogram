# Errors

## Errors from handlers

Handlers return the errors they don't handle themselves:

```go
bot.OnCommand("stats", func(ctx context.Context, u *gogram.Update) error {
	stats, err := loadStats(ctx)
	if err != nil {
		return fmt.Errorf("load stats: %w", err)
	}
	_, err = u.Reply(ctx, telegram.Text(stats))
	return err
})
```

`StartPolling` and `WebhookHandler` pass those errors to the bot's error handler. A handler that panics is recovered, and the panic reaches the error handler as a `*gogram.PanicError` holding the panic value and the stack trace; the bot keeps running.

The default error handler logs the error with the bot's logger, including the stack trace of panics. Set your own to report errors elsewhere, or to tell the user:

```go
bot := gogram.New(token, gogram.WithErrorHandler(func(ctx context.Context, u *gogram.Update, err error) {
	if p, ok := errors.AsType[*gogram.PanicError](err); ok {
		sentry.CaptureException(p) // p.Stack holds the stack trace
	}
	slog.ErrorContext(ctx, "handler failed", "update_id", u.Raw().UpdateID, "error", err)
	u.Reply(ctx, telegram.Text("Something went wrong, please try again."))
}))
```

`HandleUpdate` returns the handler's error (or a `*PanicError`) instead of calling the error handler.

## API errors

A failed API call returns one of these:

| Error | Cause |
|---|---|
| `*telegram.Error` | Telegram received the call and rejected it |
| an error wrapping `context.Canceled` or `context.DeadlineExceeded` | the call's context ended |
| any other error | a network failure, an invalid response, or an invalid request caught before sending (such as too much callback data, or a missing file) |

All of them start with `telegram:` and name the API method. They never contain the bot token.

```go
type Error struct {
	Method          string        // e.g. "sendMessage"
	Code            int           // 400, 403, 429, ...
	Description     string        // Telegram's message
	RetryAfter      time.Duration // set on rate limits
	MigrateToChatID int64         // set when a group became a supergroup
}
```

Use `errors.As` (or `errors.AsType`) to inspect it:

```go
_, err := u.Reply(ctx, telegram.Text("hi"))

if apiErr, ok := errors.AsType[*telegram.Error](err); ok {
	switch {
	case apiErr.Code == 403:
		// The user blocked the bot, or the bot was removed from the chat.
		markInactive(u.ChatID())
	case apiErr.RetryAfter > 0:
		// Rate limited.
		time.Sleep(apiErr.RetryAfter)
	case apiErr.MigrateToChatID != 0:
		// The group was upgraded; the chat has a new ID.
		updateChatID(u.ChatID(), apiErr.MigrateToChatID)
	}
}
```

Common codes:

| Code | Usual meaning |
|---|---|
| 400 | Bad request: wrong parameters, message not found, message not modified, text too long |
| 401 | Invalid token |
| 403 | The bot can't message this user or chat |
| 409 | Conflict: another instance is polling, or a webhook is set while polling |
| 429 | Too many requests; wait `RetryAfter` |

## Rate limits

Telegram allows about 30 messages per second overall, 1 per second in a single private chat, and 20 per minute in a group. gogram doesn't throttle or retry sends for you. For broadcasts, pace your sends and honour `RetryAfter`.

Polling is the exception: `StartPolling` retries failed `getMe` and `getUpdates` calls and waits `RetryAfter` when told to.

## Framework errors

| Error | Returned by |
|---|---|
| `gogram.ErrAlreadyPolling` | `StartPolling` when the bot is already polling |
| `gogram.ErrShutdownTimeout` | `StartPolling` when handlers outlast `WithShutdownTimeout` |
| `gogram.ErrNoChat`, `ErrNoCallbackQuery`, `ErrNoInlineQuery`, `ErrNoShippingQuery`, `ErrNoPreCheckoutQuery`, `ErrNoMessage` | `Update` shortcuts used on an update that doesn't fit ([details](update.md#when-a-shortcut-doesnt-apply)) |

Compare them with `errors.Is`.
