package telegram_test

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/irisked/gogram/telegram"
)

func ExampleClient_SendMessage() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))
	ctx := context.Background()

	kb := telegram.NewInlineKeyboard(
		telegram.URLButton("Docs", "https://core.telegram.org/bots/api"),
	)
	_, err := client.SendMessage(ctx, 123456789, telegram.HTML("<b>Hello</b>"),
		telegram.WithReplyMarkup(kb),
		telegram.WithDisableNotification(),
		telegram.WithProtectContent(),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_SendPhoto() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))
	ctx := context.Background()

	msg, err := client.SendPhoto(ctx, 123456789, telegram.LocalPath("cat.jpg"),
		telegram.WithCaption(telegram.Text("A cat")),
		telegram.WithHasSpoiler(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send the same photo again without uploading it.
	fileID := msg.Photo[len(msg.Photo)-1].FileID
	client.SendPhoto(ctx, 987654321, telegram.FileID(fileID))
}

func ExampleClient_SendMediaGroup() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))

	_, err := client.SendMediaGroup(context.Background(), 123456789, []telegram.InputMedia{
		&telegram.InputMediaPhoto{Media: telegram.FileURL("https://example.com/a.jpg"), Caption: "Album"},
		&telegram.InputMediaPhoto{Media: telegram.FileURL("https://example.com/b.jpg")},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_GetChatMember() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))

	member, err := client.GetChatMember(context.Background(), -1001234567890, 123456789)
	if err != nil {
		log.Fatal(err)
	}
	switch m := member.(type) {
	case *telegram.ChatMemberOwner, *telegram.ChatMemberAdministrator:
		log.Println("admin")
	case *telegram.ChatMemberRestricted:
		log.Println("restricted until", time.Unix(m.UntilDate, 0))
	case *telegram.ChatMemberBanned:
		log.Println("banned")
	}
}

// Sending to a public channel by its username.
func ExampleChatUsername() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))

	req := telegram.NewSendMessage(0, telegram.Text("New release is out!"))
	req.ChatID = telegram.ChatUsername("mychannel")
	if err := client.Do(context.Background(), req, nil); err != nil {
		log.Fatal(err)
	}
}

func ExampleError() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))

	_, err := client.SendMessage(context.Background(), 123456789, telegram.Text("hi"))
	var apiErr *telegram.Error
	if errors.As(err, &apiErr) && apiErr.RetryAfter > 0 {
		time.Sleep(apiErr.RetryAfter)
		// ...and retry.
	}
}

// Calling a method this package doesn't cover yet: any struct with JSON
// tags and a Method method is a request.
func ExampleClient_Do() {
	client := telegram.New(os.Getenv("BOT_TOKEN"))

	err := client.Do(context.Background(), &futureMethod{ChatID: 123456789}, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type futureMethod struct {
	ChatID int64 `json:"chat_id"`
}

func (*futureMethod) Method() string { return "futureMethod" }
