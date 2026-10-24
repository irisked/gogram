# Getting started

## Install

```
go get github.com/irisked/gogram
```

gogram needs Go 1.27.3 or newer and has no dependencies outside the standard library.

## Get a token

Talk to [@BotFather](https://t.me/BotFather) in Telegram, send `/newbot`, and copy the token it gives you. Keep it out of your source code; the examples read it from the `BOT_TOKEN` environment variable.

## An echo bot

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
		_, err := u.Reply(ctx, telegram.Text("Hi! Send me anything and I'll send it back."))
		return err
	})

	bot.OnText(func(ctx context.Context, u *gogram.Update) error {
		_, err := u.Reply(ctx, telegram.Text(u.Message().Text))
		return err
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := bot.StartPolling(ctx); err != nil {
		log.Fatal(err)
	}
}
```

Run it with `BOT_TOKEN=... go run .`, open your bot in Telegram and send `/start`.

What happens here:

1. `gogram.New` creates a `*gogram.Bot` and its `*telegram.Client`.
2. `OnStart` and `OnText` register handlers. A handler is a `func(ctx context.Context, u *gogram.Update) error`.
3. `StartPolling` asks Telegram for updates in a loop and runs the matching handler for each one. It blocks until the context is cancelled; here that happens on Ctrl+C.

A handler returns the errors it doesn't handle itself. By default the bot logs them; see [Errors](errors.md).

## Formatted text

Message text is a `telegram.TextData`: the text plus its parse mode.

```go
u.Reply(ctx, telegram.Text("plain text"))
u.Reply(ctx, telegram.HTML("<b>bold</b> and <i>italic</i>"))
u.Reply(ctx, telegram.MarkdownV2("*bold* and _italic_"))
u.Reply(ctx, telegram.Markdown("*bold*")) // legacy Markdown
```

Captions and other formatted parameters take a `TextData` too: `telegram.WithCaption(telegram.HTML("<i>a cat</i>"))`.

## Next

- [Routing](routing.md) explains the `On*` methods.
- [The update](update.md) lists what a handler can do.
- [Options](options.md) shows how to pass optional parameters.
