# Options

Bot API methods have a few required parameters and many optional ones. In gogram, required parameters are function arguments and optional ones are options:

```go
u.Reply(ctx, telegram.Text("hi"),
	telegram.WithReplyMarkup(keyboard),
	telegram.WithReplyToMessageID(messageID),
	telegram.WithDisableNotification(),
)
```

## Naming

The option for a parameter is `With` followed by the parameter's name: `reply_markup` is `WithReplyMarkup`, `message_thread_id` is `WithMessageThreadID`, `has_spoiler` is `WithHasSpoiler`.

- Parameters that switch something on (`disable_notification`, `has_spoiler`, `show_alert`, ...) take no argument: `telegram.WithDisableNotification()`.
- List parameters take their items as arguments: `telegram.WithEntities(e1, e2)`, `telegram.WithAllowedUpdates(telegram.UpdateMessage, telegram.UpdateCallbackQuery)`.
- Formatted text parameters take a `telegram.TextData`, which sets the parse mode too: `telegram.WithCaption(telegram.HTML("<b>x</b>"))`.
- Numbers keep the API's units: durations are in seconds, dates are Unix times (`time.Now().Add(time.Hour).Unix()`).

One option works with every method that has the parameter. Using it with a method that doesn't is a compile error, not a silently ignored parameter:

```go
// Does not compile: sendPhoto has no thumbnail parameter.
u.ReplyWithPhoto(ctx, file, telegram.WithThumbnail(thumb))
```

Each method's option type lists the options it accepts:

```
go doc github.com/irisked/gogram/telegram SendVideoOption
```

and each option's type lists the methods that accept it:

```
go doc github.com/irisked/gogram/telegram ReplyMarkupOption
```

## Shortcuts

A few options cover common cases of a more general parameter:

| Option | Same as |
|---|---|
| `WithReplyToMessageID(id)` | `WithReplyParameters(&telegram.ReplyParameters{MessageID: id})` |
| `WithDisableWebPagePreview()` | `WithLinkPreviewOptions(&telegram.LinkPreviewOptions{IsDisabled: true})` |
| `WithInlineKeyboard(kb)` | `WithReplyMarkup(kb)`, and also works with the methods that accept only inline keyboards, such as the edit methods |

## Common options

Options available on most send methods:

| Option | Effect |
|---|---|
| `WithReplyMarkup(markup)` | attach a keyboard, remove one, or force a reply ([Keyboards](keyboards.md)) |
| `WithReplyToMessageID(id)` | reply to a message in the same chat |
| `WithMessageThreadID(id)` | send to a forum topic |
| `WithDisableNotification()` | send silently |
| `WithProtectContent()` | prevent forwarding and saving |
| `WithBusinessConnectionID(id)` | send on behalf of a connected business account |

Text messages (`SendMessage`, `EditMessageText`):

| Option | Effect |
|---|---|
| `WithDisableWebPagePreview()` | no link preview |
| `WithLinkPreviewOptions(&telegram.LinkPreviewOptions{...})` | fine-grained link preview control |
| `WithEntities(entities...)` | formatting as entities instead of a parse mode |

Media (`SendPhoto`, `SendVideo`, `SendAudio`, ...):

| Option | Effect |
|---|---|
| `WithCaption(text)` | caption with parse mode, e.g. `telegram.HTML("<b>x</b>")` |
| `WithShowCaptionAboveMedia()` | caption above the media |
| `WithHasSpoiler()` | blur the media until tapped |
| `WithThumbnail(file)` | custom thumbnail |
| `WithDuration(seconds)`, `WithWidth(w)`, `WithHeight(h)` | media metadata |
| `WithSupportsStreaming()` | video can be streamed |
| `WithPerformer(s)`, `WithTitle(s)` | audio metadata |

## Building requests yourself

Each method also has a request type named like the method, with a constructor that takes the required parameters. The fields are exported, so you can set them directly:

```go
req := telegram.NewSendMessage(chatID, telegram.Text("hi"))
req.ProtectContent = true
req.ReplyParameters = &telegram.ReplyParameters{MessageID: id, AllowSendingWithoutReply: true}

var msg telegram.Message
err := client.Do(ctx, req, &msg)
```

This is handy when you build a request in several steps, or to send to a chat by username. See [Using the client directly](client.md).
