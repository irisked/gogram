package gogram

import "testing"

func TestTreeInsertAndGet(t *testing.T) {
	tr := newTree()
	h1 := Handler(func(ctx *Context) {})
	h2 := Handler(func(ctx *Context) {})

	if _, updated := tr.Insert("/message", h1); updated {
		t.Fatalf("expected first insert to not report an update")
	}
	if _, found := tr.Get("/message"); !found {
		t.Fatalf("expected to find /message")
	}
	if _, updated := tr.Insert("/message", h2); !updated {
		t.Fatalf("expected second insert of the same key to report an update")
	}
}

func TestTreeLongestPrefix(t *testing.T) {
	tr := newTree()
	tr.Insert("/message", Handler(func(ctx *Context) {}))
	tr.Insert("/message/text", Handler(func(ctx *Context) {}))

	if key, _, ok := tr.LongestPrefix("/message/text/extra"); !ok || key != "/message/text" {
		t.Fatalf("expected longest prefix /message/text, got %q (ok=%v)", key, ok)
	}
	if key, _, ok := tr.LongestPrefix("/message/other"); !ok || key != "/message" {
		t.Fatalf("expected longest prefix /message, got %q (ok=%v)", key, ok)
	}
}

func TestTreeNoMatch(t *testing.T) {
	tr := newTree()
	if _, _, ok := tr.LongestPrefix("/anything"); ok {
		t.Fatalf("expected no match on an empty tree")
	}
}
