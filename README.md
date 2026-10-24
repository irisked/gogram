# gogram

A Telegram bot framework for Go, with a complete Bot API client generated from the API reference.

```
go get github.com/irisked/gogram
```

- `gogram`: the bot framework. It handles routing, middleware, long polling and webhooks.
- `gogram/telegram`: a standalone Bot API client with every method and type of the Bot API, keyboards and file uploads.
- `gogram/sandbox`: a fake Bot API server and update builders for testing bots.

Full guides are in [docs/](docs/README.md).

## Quick start

```go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/irisked/gogram"
	"github.com/irisked/gogram/telegram"
)

func main() {
	bot := gogram.New(os.Getenv("BOT_TOKEN"))

	bot.OnStart(func(ctx context.Context, u *gogram.Update) error {
		_, err := u.Reply(ctx, telegram.Text("Hi! Send me anything."))
		return err
	})
	bot.OnText(func(ctx context.Context, u *gogram.Update) error {
		_, err := u.Reply(ctx, telegram.Text(u.Message().Text))
		return err
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := bot.StartPolling(ctx); err != nil { // blocks until Ctrl+C
		log.Fatal(err)
	}
}
```

## Handlers and routing

Each update goes to the most specific handler registered for it. For example:

- `/start` goes to `OnStart`, or else `OnMessage`, or else `OnAny`.
- A photo goes to `OnPhoto`, or else `OnMessage`, or else `OnAny`.
- A button press goes to `OnCallback(route)`, or else `OnCallbackQuery`, or else `OnAny`.

```go
bot.OnCommand("settings", handler)     // /settings, /Settings, /settings@my_bot
bot.OnPhoto(handler)
bot.OnMessage(handler)                 // any other message
bot.OnPollAnswer(handler)
bot.OnUpdate(telegram.UpdateChatBoost, handler)
bot.OnAny(handler)                     // everything else
```

In groups, commands addressed to another bot (`/start@other_bot`) are routed as text, so your bot doesn't answer them.

Handlers get a `context.Context` and a `*gogram.Update`, the incoming update plus shortcuts that act on it, and return an error:

```go
bot.OnCommand("ban", func(ctx context.Context, u *gogram.Update) error {
	args := u.Args() // "/ban 42" -> ["42"]
	// u.From(), u.Chat(), u.ChatID(), u.Message(), u.CallbackQuery() ...
	_, err := u.Reply(ctx, telegram.HTML("<b>banned "+args[0]+"</b>"))
	return err
})
```

Returned errors and recovered panics, with their stack trace, go to the error handler, which logs them by default:

```go
bot := gogram.New(token, gogram.WithErrorHandler(func(ctx context.Context, u *gogram.Update, err error) {
	// report err
}))
```

Updates from the same chat are handled one at a time and in order; different chats are handled concurrently. `gogram.WithConcurrency(n)` caps the total.

## Options

Optional API parameters are options named after them. Using one with a method that doesn't have that parameter is a compile error:

```go
u.Reply(ctx, telegram.Text("hi"),
	telegram.WithReplyMarkup(kb),
	telegram.WithReplyToMessageID(id),
	telegram.WithDisableNotification(),
)
```

## Keyboards and callbacks

```go
bot.OnCommand("vote", func(ctx context.Context, u *gogram.Update) error {
	kb := telegram.NewInlineKeyboard(
		telegram.CallbackButton("👍", "vote", "up"),
		telegram.CallbackButton("👎", "vote", "down"),
	).Row(
		telegram.URLButton("Rules", "https://example.com/rules"),
	)
	_, err := u.Reply(ctx, telegram.Text("Do you like it?"), telegram.WithReplyMarkup(kb))
	return err
})

bot.OnCallback("vote", func(ctx context.Context, u *gogram.Update) error {
	if err := u.AnswerCallbackQuery(ctx, telegram.WithText("Thanks!")); err != nil {
		return err
	}
	_, err := u.EditMessageText(ctx, telegram.Text("You voted "+u.Args()[0]))
	return err
})
```

Reply keyboards work the same way: `telegram.NewReplyKeyboard(telegram.ContactButton("Share phone"))`.

## Files

```go
u.ReplyWithPhoto(ctx, telegram.LocalPath("cat.jpg"))           // upload
u.ReplyWithPhoto(ctx, telegram.FileID(fileID))                 // reuse a file Telegram has
u.ReplyWithPhoto(ctx, telegram.FileURL("https://..."))         // let Telegram download it
u.ReplyWithPhoto(ctx, telegram.FileReader("chart.png", data))  // upload from memory

u.ReplyWithMediaGroup(ctx, []telegram.InputMedia{
	&telegram.InputMediaPhoto{Media: telegram.FileID(a), Caption: "album"},
	&telegram.InputMediaPhoto{Media: telegram.LocalPath("b.jpg")},
})
```

Requests with uploads are sent as `multipart/form-data` automatically.

## Middleware

```go
bot.Use(func(next gogram.Handler) gogram.Handler {
	return func(ctx context.Context, u *gogram.Update) error {
		if u.From() == nil || u.From().ID != adminID {
			return nil
		}
		return next(ctx, u)
	}
})
```

## Webhooks

```go
bot.Client().SetWebhook(ctx, "https://example.com/bot", telegram.WithSecretToken(secret))
http.Handle("/bot", bot.WebhookHandler(secret))
```

## Errors

API failures are returned as `*telegram.Error`:

```go
if apiErr, ok := errors.AsType[*telegram.Error](err); ok && apiErr.RetryAfter > 0 {
	time.Sleep(apiErr.RetryAfter)
}
```

## Testing

```go
srv := sandbox.NewServer(t)
bot := srv.Bot()
registerHandlers(bot)

bot.HandleUpdate(t.Context(), sandbox.Command("/start"))
sent := srv.Calls("sendMessage") // what the bot sent
```

## Using the client on its own

```go
client := telegram.New(token)
me, err := client.GetMe(ctx)
```

`client.Do(ctx, req, &result)` sends any request, including your own types for Bot API methods newer than the package.

## Bot API version

The `telegram` package is generated from the Bot API reference at core.telegram.org, via a checksummed snapshot of it in `internal/apispec`. The version and source are named at the top of every `*_gen.go` file. Moving to a new Bot API version is `make update-api`; see [Updating the Bot API](docs/generating.md).

## Development

```
make            # list the tasks
make check      # lint, test, and check the generated code is up to date
make update-api # move to the newest Bot API version
```
