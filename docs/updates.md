# Polling and webhooks

A bot receives updates in one of two ways:

- **Long polling**: the bot repeatedly asks Telegram for new updates. It works anywhere, including behind NAT and on a laptop, and needs no public URL.
- **Webhooks**: Telegram sends each update to your HTTPS endpoint. This suits servers that already handle HTTP and serverless platforms.

Only one can be active at a time: while a webhook is set, polling fails.

## Long polling

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

if err := bot.StartPolling(ctx); err != nil {
	log.Fatal(err)
}
```

`StartPolling` blocks until `ctx` is cancelled, then returns nil. It returns `gogram.ErrAlreadyPolling` if the bot is already polling.

While it runs:

- It first looks up the bot's username with `getMe`, unless you gave it with `WithUsername` (see [Routing](routing.md#commands)).
- Each request waits up to the poll timeout (30 seconds by default) for new updates, so updates arrive immediately without busy polling.
- A failed call is logged and retried after 1 second, doubling up to 30 seconds. On a rate limit it waits as long as Telegram asks.

### Order and concurrency

Updates from the same chat are handled one at a time, in the order Telegram sent them, so a conversation is never processed out of order. Updates from different chats are handled concurrently. Updates without a chat, such as inline queries, are ordered per user.

`WithConcurrency(n)` limits how many updates are in progress at once, counting those waiting behind an earlier update from their chat. When the limit is reached, polling waits until an update finishes, so a burst of updates can't use unbounded memory or goroutines. The default is no limit.

```go
bot := gogram.New(token, gogram.WithConcurrency(100))
```

### Shutting down

When `ctx` is cancelled, `StartPolling`:

1. stops asking for updates,
2. waits for running and queued handlers to return,
3. tells Telegram which updates were handled, so they aren't delivered again on the next start,
4. returns nil.

Running handlers aren't interrupted: their context keeps the values of `ctx` but isn't cancelled with it, so replies sent during shutdown still go out.

A handler that never returns would block shutdown forever. `WithShutdownTimeout(d)` bounds the wait: when it expires, the handlers' context is cancelled and `StartPolling` returns `gogram.ErrShutdownTimeout`.

```go
bot := gogram.New(token, gogram.WithShutdownTimeout(10*time.Second))
```

### Polling options

```go
bot := gogram.New(token,
	gogram.WithPollTimeout(50*time.Second),
	gogram.WithAllowedUpdates(
		telegram.UpdateMessage,
		telegram.UpdateCallbackQuery,
		telegram.UpdateChatMember,
	),
)
```

| Option | Effect |
|---|---|
| `WithPollTimeout(d)` | how long each `getUpdates` waits for updates (default 30s) |
| `WithAllowedUpdates(types...)` | which update types Telegram sends; `telegram.AllUpdateTypes` lists them all |
| `WithConcurrency(n)` | at most n updates in progress (default: no limit) |
| `WithShutdownTimeout(d)` | how long shutdown waits for handlers (default: as long as they take) |
| `WithUsername(name)` | the bot's username, instead of looking it up |
| `WithErrorHandler(h)` | what to do with handler errors ([Errors](errors.md)) |
| `WithLogger(l)` | `*slog.Logger` for polling errors and the default error handler (default `slog.Default()`) |

About `WithAllowedUpdates`: Telegram remembers the last setting it received, from `getUpdates` or `setWebhook`. Without the option, the bot sends none and the remembered setting applies; for a new bot that is every type except `chat_member`, `message_reaction` and `message_reaction_count`. Pass the option to be sure what you get.

To use a preconfigured client, create the bot with `gogram.NewWithClient(client, opts...)` ([Using the client directly](client.md)). If you give that client an `http.Client` with a `Timeout`, make it longer than the poll timeout, or every poll fails.

## Webhooks

`Bot.WebhookHandler` returns an `http.Handler` that decodes each update and handles it before responding:

```go
const secret = "a-long-random-string"

func main() {
	bot := gogram.New(os.Getenv("BOT_TOKEN"))
	// ... register handlers ...

	err := bot.Client().SetWebhook(context.Background(), "https://bot.example.com/telegram",
		telegram.WithSecretToken(secret),
		telegram.WithAllowedUpdates(telegram.UpdateMessage, telegram.UpdateCallbackQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/telegram", bot.WebhookHandler(secret))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

- Telegram only calls HTTPS URLs on ports 443, 80, 88 or 8443. Terminate TLS in front of the app or use `http.ListenAndServeTLS`. For a self-signed certificate, upload it with `WithCertificate`.
- With a secret token, the handler rejects requests that don't carry it in the `X-Telegram-Bot-Api-Secret-Token` header. Always set one; otherwise anyone who finds the URL can send fake updates.
- Updates are handled synchronously, and Telegram waits for the response before sending more to that connection. `WithMaxConnections` controls how many arrive in parallel (default 40). Telegram orders updates per chat itself.
- The handler's context is the HTTP request's context.
- A failed handler's error goes to the error handler, and the response is still a success: answering with an error makes Telegram send the update again.
- The first update looks up the bot's username with `getMe`, unless `WithUsername` gave it.

`SetWebhook` only needs to be called once, or whenever the URL changes; Telegram remembers it.

### Switching back to polling

```go
bot.Client().DeleteWebhook(ctx) // add telegram.WithDropPendingUpdates() to discard queued updates
```

## Other update sources

`Bot.HandleUpdate(ctx, update)` routes one update, waits for its handler and returns its error. Use it to feed updates from a queue, a custom HTTP framework, or a test:

```go
func (s *Server) telegram(w http.ResponseWriter, r *http.Request) {
	var u telegram.Update
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.bot.HandleUpdate(r.Context(), u); err != nil {
		slog.ErrorContext(r.Context(), "update failed", "error", err)
	}
}
```

`HandleUpdate` doesn't call the error handler, and doesn't look up the bot's username; call `bot.Init(ctx)` once at startup, or use `WithUsername`.
