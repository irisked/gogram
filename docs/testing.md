# Testing

Package `sandbox` tests bots without talking to Telegram. It has a fake Bot API server that records what the bot sends and answers like Telegram, and builders for incoming updates.

## A first test

`HandleUpdate` routes an update, waits for its handler and returns its error, so assertions can follow it directly. A table test covers each command with one row:

```go
func TestCommands(t *testing.T) {
	tests := []struct {
		name   string
		update telegram.Update
		want   string // text of the reply
	}{
		{"start", sandbox.Command("/start"), "Welcome!"},
		{"help", sandbox.Command("/help"), "Send me a photo."},
		{"unknown command", sandbox.Command("/nope"), "I don't know that command."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			bot := srv.Bot()
			registerHandlers(bot) // the same function main uses

			if err := bot.HandleUpdate(t.Context(), tt.update); err != nil {
				t.Fatal(err)
			}

			sent := srv.Calls("sendMessage")
			if len(sent) != 1 || sent[0].String("text") != tt.want {
				t.Fatalf("sent %v, want one reply %q", sent, tt.want)
			}
		})
	}
}
```

`srv.Bot(opts...)` returns a bot that talks to the server and already knows its username (`sandbox.BotUsername`). `srv.Client()` returns just the client.

## Updates

| Builder | Update |
|---|---|
| `Message(text)` | a text message from `sandbox.User` in their private chat; text starting with `/` is marked as a command |
| `Command(text)` | the same, adding the `/` if missing |
| `GroupMessage(text)` | a text message from the user in a group (`sandbox.GroupID`) |
| `Callback(data)` | a press of a button with that callback data, on a bot message in the private chat |
| `InlineQuery(query)` | an inline query from the user |

They return a plain `telegram.Update`, so change any field you need, or build other updates yourself. To test with recorded real updates, unmarshal their JSON into `telegram.Update`.

## What the bot sent

| Method | Returns |
|---|---|
| `srv.Requests()` | every request received, except `getUpdates` |
| `srv.Calls(method)` | the requests for one method |
| `srv.WaitFor(method, n)` | waits until `n` requests for a method have arrived, for asynchronous code such as polling |

A `sandbox.Request` has the `Method`, the `Params` and any uploaded `Files`. `r.String(name)` and `r.Int(name)` read a parameter; objects such as `reply_markup` come back as JSON, which `r.String` returns.

## How the server answers

By default the server answers:

- `getMe` with the test bot;
- `getUpdates` with the updates queued by `srv.AddUpdates`;
- the send, copy, forward and edit methods with a message built from the request (its chat, text and caption);
- every other method with an empty result of the right type.

Change the answer of a method with `Handle`. Return a `*telegram.Error` to answer with an API error:

```go
srv.Handle("getChatMember", func(r sandbox.Request) (any, error) {
	return map[string]any{"status": "administrator", "user": map[string]any{"id": r.Int("user_id")}}, nil
})

srv.Handle("sendMessage", func(sandbox.Request) (any, error) {
	return nil, &telegram.Error{Code: 403, Description: "Forbidden: bot was blocked by the user"}
})
```

## Testing polling

Queue updates and wait for the requests they cause:

```go
srv := sandbox.NewServer(t)
srv.AddUpdates(sandbox.Message("one"), sandbox.Message("two"))
bot := srv.Bot()
registerHandlers(bot)

ctx, cancel := context.WithCancel(t.Context())
done := make(chan error)
go func() { done <- bot.StartPolling(ctx) }()

replies := srv.WaitFor("sendMessage", 2)
cancel()
if err := <-done; err != nil {
	t.Fatal(err)
}
```

## Testing a handler alone

`gogram.NewUpdate` builds the `*gogram.Update` a handler receives, without a bot or routing:

```go
u := gogram.NewUpdate(srv.Client(), sandbox.Command("/start"))
err := startHandler(t.Context(), u)
```

Handlers that only read the update don't need a client at all; pass nil.

## Testing webhooks

`WebhookHandler` is a normal `http.Handler`:

```go
h := bot.WebhookHandler("secret")
body, _ := json.Marshal(sandbox.Message("hi"))

req := httptest.NewRequest(http.MethodPost, "/telegram", bytes.NewReader(body))
req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "secret")
rec := httptest.NewRecorder()
h.ServeHTTP(rec, req)
```
