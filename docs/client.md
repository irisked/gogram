# Using the client directly

`telegram.Client` is a complete Bot API client. `gogram` uses it internally, but it doesn't depend on the framework: use it on its own for scripts, cron jobs, or services that only send messages.

```go
client := telegram.New(os.Getenv("BOT_TOKEN"))

me, err := client.GetMe(ctx)
if err != nil {
	log.Fatal(err)
}
fmt.Println("running as", me.Username)

_, err = client.SendMessage(ctx, chatID, telegram.Text("Deploy finished ✅"))
```

Inside a handler, the bot's client is `u.Client()`; elsewhere it's `bot.Client()`.

## Coverage

The package is generated from the Bot API reference (see [Updating the Bot API](generating.md)), so it has every method and type of the Bot API version it was generated from. The version is named at the top of the generated files.

## Method conventions

Every Bot API method is a `Client` method with the same name in Go style (`sendMessage` → `SendMessage`):

- The first argument is a `context.Context`; cancelling it aborts the call.
- Required parameters come next, in the order the API reference lists them.
- Optional parameters are [options](options.md).
- Formatted text parameters, such as `text` with its `parse_mode`, are one `telegram.TextData`.
- Edit methods take a `telegram.MessageTarget` first: `telegram.ChatMessage(chatID, messageID)` or `telegram.InlineMessage(inlineMessageID)`.
- Methods that return an object return a pointer to it (`*telegram.Message`). Methods that return a list return a slice.
- Methods whose API result is only `true` return just an `error`.
- Edit methods return the edited message, or nil when you edited an inline message, since Telegram returns no message for those.

Each method's doc comment includes the API reference text and links to its section.

## Interface types

Some objects come in several kinds: a `ChatMember` is an owner, an administrator, a member, a restricted member, a former member or a banned user; an `InputMedia` is a photo, a video, an audio file and so on. Each is an interface implemented by one struct per kind. Use a type switch:

```go
member, err := client.GetChatMember(ctx, chatID, userID)
if err != nil {
	return err
}
switch m := member.(type) {
case *telegram.ChatMemberOwner, *telegram.ChatMemberAdministrator:
	// an admin
case *telegram.ChatMemberRestricted:
	until := time.Unix(m.UntilDate, 0)
case *telegram.ChatMemberBanned:
	// banned
}
```

When sending, pass a pointer to the kind you mean: `&telegram.ReactionTypeEmoji{Emoji: "👍"}`. The field that tells kinds apart on the wire (`type`, `status` or `source`) isn't part of the structs; it is added and read for you.

A kind added in a newer Bot API version decodes as `*telegram.UnknownVariant`, which holds the raw JSON, instead of failing.

## Configuration

```go
client := telegram.New(token,
	telegram.WithHTTPClient(&http.Client{Transport: myTransport}),
	telegram.WithBaseURL("http://localhost:8081"),
)
bot := gogram.NewWithClient(client)
```

| Option | Effect |
|---|---|
| `WithHTTPClient(hc)` | the `*http.Client` to use, e.g. for proxies or instrumentation. Default: `http.DefaultClient` |
| `WithBaseURL(url)` | another Bot API server, such as a [self-hosted one](https://github.com/tdlib/telegram-bot-api) with higher file limits, or a test server |

Prefer per-call deadlines through `context` over `http.Client.Timeout`: a client-wide timeout shorter than the polling timeout breaks long polling.

A `Client` is safe for concurrent use. Create one per token and share it.

## Request types

Each method has a request type named like the method (`SendMessage`), with a constructor taking the required parameters (`NewSendMessage`). All of them implement `telegram.Request`:

```go
type Request interface {
	Method() string // e.g. "sendMessage"
}
```

`Client.Do` sends any request and decodes its result:

```go
req := telegram.NewSendMessage(chatID, telegram.Text("hi"))
req.ReplyParameters = &telegram.ReplyParameters{MessageID: id, AllowSendingWithoutReply: true}

var msg telegram.Message
if err := client.Do(ctx, req, &msg); err != nil {
	return err
}
```

Pass `nil` as the result to discard it.

### Sending to a chat by username

Public channels and supergroups can be addressed by username. Client methods take numeric chat IDs, so set the request's `ChatID` and send it with `Do`:

```go
req := telegram.NewSendMessage(0, telegram.Text("New release is out!"))
req.ChatID = telegram.ChatUsername("mychannel")
err := client.Do(ctx, req, nil)
```

## Calling methods gogram doesn't cover

Any struct with JSON tags and a `Method` method is a request. That covers Bot API methods newer than the generated package:

```go
type FutureMethod struct {
	ChatID int64  `json:"chat_id"`
	Value  string `json:"value,omitempty"`
}

func (*FutureMethod) Method() string { return "futureMethod" }

err := client.Do(ctx, &FutureMethod{ChatID: chatID, Value: "x"}, nil)
```

`*telegram.InputFile` fields work in your own requests too: uploads are detected and sent as multipart automatically.
