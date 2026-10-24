// Package gogram is a framework for Telegram bots.
//
// A Bot routes incoming updates to handlers registered with the On*
// methods. Each handler receives a context.Context and an *Update that
// wraps the raw update and offers shortcuts such as Reply, and returns an
// error. The Bot API itself lives in the telegram subpackage: its Client
// is available from Bot.Client and Update.Client.
//
//	bot := gogram.New(os.Getenv("BOT_TOKEN"))
//
//	bot.OnStart(func(ctx context.Context, u *gogram.Update) error {
//		_, err := u.Reply(ctx, telegram.Text("Hello!"))
//		return err
//	})
//	bot.OnText(func(ctx context.Context, u *gogram.Update) error {
//		_, err := u.Reply(ctx, telegram.Text(u.Message().Text))
//		return err
//	})
//
//	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
//	defer stop()
//	if err := bot.StartPolling(ctx); err != nil {
//		log.Fatal(err)
//	}
//
// # Routing
//
// Every update has a route made of its update type, then for messages
// the message type and command, and for callback queries the callback
// route set by telegram.CallbackButton. An update goes to the most
// specific handler registered for its route; for example a "/start"
// command goes to the OnStart handler if there is one, otherwise to
// OnMessage, otherwise to OnAny. Updates with no matching handler are
// dropped. Commands addressed to another bot ("/start@other_bot") are
// routed as text.
//
// # Errors
//
// Errors returned by handlers, and recovered panics as a *PanicError, go
// to the ErrorHandler set with WithErrorHandler. The default logs them.
//
// # Concurrency
//
// While polling, updates from the same chat are handled one at a time,
// in order, and updates from different chats concurrently, so handlers
// that share state must synchronize access to it. WithConcurrency caps
// how many updates are in progress.
//
// # Testing
//
// Package sandbox provides a fake Bot API server and update builders.
package gogram
