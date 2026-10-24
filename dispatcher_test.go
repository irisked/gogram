package gogram

import (
	"testing"

	"github.com/irisked/gogram/types"
)

func newTestContext(u types.Update) *Context {
	return &Context{update: u, State: make(map[string]interface{})}
}

func TestDispatcherOnTextDispatchesToHandler(t *testing.T) {
	d := newDispatcher()
	called := false
	d.OnText(func(ctx *Context) { called = true })

	u := types.Update{Message: &types.Message{Text: "hello"}}
	d.Handle(newTestContext(u))

	if !called {
		t.Fatalf("expected OnText handler to be called")
	}
}

func TestDispatcherOnCommandDispatchesToMatchingHandler(t *testing.T) {
	d := newDispatcher()
	var got string
	d.OnCommand("start", func(ctx *Context) { got = "start" })
	d.OnCommand("help", func(ctx *Context) { got = "help" })

	entities := []types.MessageEntity{{Type: "bot_command", Offset: 0, Length: 6}}
	u := types.Update{Message: &types.Message{Text: "/start", Entities: &entities}}
	d.Handle(newTestContext(u))

	if got != "start" {
		t.Fatalf("expected the start handler to be called, got %q", got)
	}
}

func TestDispatcherNoHandlerRegisteredDoesNotPanic(t *testing.T) {
	d := newDispatcher()
	u := types.Update{Message: &types.Message{Text: "hello"}}
	d.Handle(newTestContext(u))
}
