# gogram documentation

gogram is a Telegram bot framework for Go. It has three public packages:

| Package | Import path | What it is |
|---|---|---|
| `gogram` | `github.com/irisked/gogram` | The bot framework: routing, handlers, middleware, long polling, webhooks. |
| `telegram` | `github.com/irisked/gogram/telegram` | A complete Bot API client, generated from the API reference: every method and type, options, keyboards, file uploads. |
| `sandbox` | `github.com/irisked/gogram/sandbox` | A fake Bot API server and update builders for testing bots. |

Most bots use `gogram` to receive updates and `telegram` for the values passed to and returned from the API. The `telegram` package also works on its own when you only need to call the API.

## Guides

1. [Getting started](getting-started.md): install the module and write your first bot.
2. [Routing](routing.md): which handler receives which update.
3. [The update](update.md): reading the update and replying.
4. [Options](options.md): passing optional API parameters.
5. [Keyboards and callbacks](keyboards.md): inline and reply keyboards, button presses.
6. [Sending files](files.md): uploads, file IDs, URLs, albums, downloads.
7. [Middleware](middleware.md): code that runs around every handler.
8. [Polling and webhooks](updates.md): receiving updates, ordering, concurrency and shutdown.
9. [Using the client directly](client.md): the `telegram.Client`, requests, interface types, configuration.
10. [Errors](errors.md): handler errors, API errors, rate limits.
11. [Testing](testing.md): testing bots with `sandbox`.
12. [Updating the Bot API](generating.md): how the `telegram` package is generated, and the Make targets.

## Reference

The full API reference is in the Go doc comments:

```
go doc -all github.com/irisked/gogram
go doc github.com/irisked/gogram/telegram Client.SendMessage
go doc github.com/irisked/gogram/telegram SendMessageOption
```

Every `telegram.Client` method links to the matching section of the [Telegram Bot API reference](https://core.telegram.org/bots/api).
