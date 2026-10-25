# Sending files

Methods that send files take a `*telegram.InputFile`. There are several ways to make one:

| Constructor | Use it for |
|---|---|
| `telegram.FileID(id)` | a file Telegram already has, e.g. one you received or sent before |
| `telegram.FileURL(url)` | a file on the web; Telegram downloads it |
| `telegram.LocalPath(path)` | a file on disk, opened when sending and closed afterwards |
| `telegram.FileReader(name, r)` | content from any `io.Reader`, e.g. generated in memory |
| `telegram.LocalFile(f)` | a file you already opened; you close it |

```go
// Upload from disk.
u.ReplyWithDocument(ctx, telegram.LocalPath("report.pdf"))

// Upload from memory.
u.ReplyWithPhoto(ctx, telegram.FileReader("chart.png", bytes.NewReader(png)))

// Let Telegram fetch it.
u.ReplyWithPhoto(ctx, telegram.FileURL("https://example.com/cat.jpg"))

// Reuse a file without uploading it again.
u.ReplyWithPhoto(ctx, telegram.FileID(fileID))
```

When a request contains an upload, gogram sends it as `multipart/form-data` and streams the content. Otherwise the request is plain JSON. You don't have to choose.

Things to know about uploads:

- Every upload is opened before the request is sent, so a missing `LocalPath` file fails the call without sending anything.
- An upload from `FileReader` or `LocalFile` reads its `io.Reader` once, so it can be sent only once. One from `LocalPath` opens the file again each time.
- Telegram limits uploads to 50 MB (10 MB for photos) on the public API server. URLs are limited to 20 MB, and 5 MB for photos.

## Reusing files

Every file Telegram sends you (or that you send) has a `FileID`. Sending by ID is instant and free, so store IDs of files you send often:

```go
msg, err := u.ReplyWithPhoto(ctx, telegram.LocalPath("cat.jpg"))
if err != nil {
	return err
}
largest := msg.Photo[len(msg.Photo)-1] // photos come in several sizes
saveFileID(largest.FileID)
```

## Captions and media options

```go
u.ReplyWithVideo(ctx, telegram.FileID(id),
	telegram.WithCaption(telegram.HTML("<b>Demo</b>")),
	telegram.WithDuration(42), // seconds
	telegram.WithSupportsStreaming(),
	telegram.WithThumbnail(telegram.LocalPath("thumb.jpg")),
)
```

See [Options](options.md).

## Albums

`ReplyWithMediaGroup` (or `Client.SendMediaGroup`) sends 2–10 items as one album. Items are `telegram.InputMedia` values: `*InputMediaPhoto`, `*InputMediaVideo`, `*InputMediaAudio`, `*InputMediaDocument` and `*InputMediaLivePhoto`. Uploads, IDs and URLs can be mixed:

```go
u.ReplyWithMediaGroup(ctx, []telegram.InputMedia{
	&telegram.InputMediaPhoto{Media: telegram.FileID(first), Caption: "Our trip"},
	&telegram.InputMediaPhoto{Media: telegram.LocalPath("beach.jpg")},
	&telegram.InputMediaVideo{Media: telegram.FileURL("https://example.com/v.mp4"), SupportsStreaming: true},
})
```

Photos and videos can share an album. Audio files and documents can only be grouped with their own kind.

## Replacing media

`EditMessageMedia` swaps the media of a message, for example in a gallery driven by buttons:

```go
bot.OnCallback("next", func(ctx context.Context, u *gogram.Update) error {
	_, err := u.EditMessageMedia(ctx, &telegram.InputMediaPhoto{Media: telegram.FileID(nextID)},
		telegram.WithInlineKeyboard(galleryKeyboard))
	if err != nil {
		return err
	}
	return u.AnswerCallbackQuery(ctx)
})
```

## Downloading files

Received media has a `FileID`. `GetFile` turns it into a path, and `FileURL` into a download URL:

```go
bot.OnDocument(func(ctx context.Context, u *gogram.Update) error {
	file, err := u.Client().GetFile(ctx, u.Message().Document.FileID)
	if err != nil {
		return err
	}
	url := u.Client().FileURL(file.FilePath)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	// ...
})
```

The URL contains your bot token, so don't log it or show it to users. The public API server only serves files up to 20 MB.
