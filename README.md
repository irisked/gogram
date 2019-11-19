# gogram
Telegram bot framework.

```
  import (
    "github.com/irisked/gogram/args"
    "github.com/irisked/gogram/option"
    "github.com/irisked/gogram/telegram/keyboard/inline"
    "github.com/irisked/gogram"
  )
	bot := gogram.New("<token>")
	bot.OnMessage(func (ctx *gogram.Context) {
		upd := ctx.Update()
		ctx.Reply(args.Text("test response"))
	})
	bot.OnEditedMessage(func (ctx *gogram.Context) {
		ctx.Reply(args.Text("test edited response"))
	})
	bot.OnSticker(func (ctx *gogram.Context) {
		upd := ctx.Update()
		kb := inline.NewKeyboard(
			inline.NewButton("test", inline.Callback("test-callback", "1")),
			inline.NewButton("test2", inline.Callback("test", "1")),
			inline.NewButton("test2", inline.Callback("test-callback1", "1")),
		)
		ctx.Reply(args.Text("stiker response"), option.Keyboard(kb))
	})
	bot.OnCommand("test", func (ctx *gogram.Context) {
		ctx.Reply(args.Markdown("```test```"))
	})
	bot.OnCallbackQuery(func (ctx *gogram.Context) {
		upd := ctx.Update()
		ctx.Reply(args.Markdown("```test-generic```"))
	})
	bot.OnCallback("test-callback", func (ctx *gogram.Context) {
		upd := ctx.Update()
		ctx.Reply(args.Markdown("```test-callback```"))
	})
	bot.OnCallback("test-callback1", func (ctx *gogram.Context) {
		upd := ctx.Update()
		ctx.Reply(args.Markdown("```test-callback1```"))
	})
	bot.StartPolling()
```