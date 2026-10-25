# Middleware

Middleware is code that runs around every handler, for logging, access control, timing, rate limiting and similar concerns. A middleware takes the next handler and returns a new one:

```go
type Middleware func(next Handler) Handler
```

Register middleware with `Bot.Use`:

```go
bot.Use(logging, adminsOnly)
```

Middleware runs in the order it was added: here `logging` runs first and calls `adminsOnly`, which calls the handler. It wraps the handler chosen by [routing](routing.md), so it only runs for updates that have a handler.

## Examples

### Logging

Middleware sees the error the handler returns:

```go
func logging(next gogram.Handler) gogram.Handler {
	return func(ctx context.Context, u *gogram.Update) error {
		start := time.Now()
		err := next(ctx, u)
		slog.InfoContext(ctx, "update handled",
			"id", u.Raw().UpdateID,
			"type", u.Raw().Type(),
			"took", time.Since(start),
			"error", err)
		return err
	}
}
```

### Access control

Not calling `next` stops the update:

```go
func adminsOnly(admins ...int64) gogram.Middleware {
	return func(next gogram.Handler) gogram.Handler {
		return func(ctx context.Context, u *gogram.Update) error {
			if from := u.From(); from == nil || !slices.Contains(admins, from.ID) {
				return nil
			}
			return next(ctx, u)
		}
	}
}

bot.Use(adminsOnly(12345, 67890))
```

### Passing values to handlers

Add values to the context, with a key type of your own:

```go
type userKey struct{}

func loadUser(db *DB) gogram.Middleware {
	return func(next gogram.Handler) gogram.Handler {
		return func(ctx context.Context, u *gogram.Update) error {
			if from := u.From(); from != nil {
				user, err := db.User(ctx, from.ID)
				if err != nil {
					return err
				}
				ctx = context.WithValue(ctx, userKey{}, user)
			}
			return next(ctx, u)
		}
	}
}

bot.OnCommand("me", func(ctx context.Context, u *gogram.Update) error {
	user, _ := ctx.Value(userKey{}).(*User)
	// ...
	return nil
})
```

### Timeouts

```go
func timeout(d time.Duration) gogram.Middleware {
	return func(next gogram.Handler) gogram.Handler {
		return func(ctx context.Context, u *gogram.Update) error {
			ctx, cancel := context.WithTimeout(ctx, d)
			defer cancel()
			return next(ctx, u)
		}
	}
}
```

### "Typing..." while working

```go
func typing(next gogram.Handler) gogram.Handler {
	return func(ctx context.Context, u *gogram.Update) error {
		if u.ChatID() != 0 {
			u.SendChatAction(ctx, telegram.ChatActionTyping)
		}
		return next(ctx, u)
	}
}
```

## Applying middleware to some handlers only

`Use` applies to all handlers. To wrap one handler, call the middleware directly:

```go
bot.OnCommand("ban", adminsOnly(12345)(banHandler))
```

## Panics

The bot recovers panics from handlers and middleware and passes them to the error handler as a `*gogram.PanicError`, with the stack trace, so one bad update doesn't stop the bot. See [Errors](errors.md#errors-from-handlers).
