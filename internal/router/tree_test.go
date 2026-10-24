package router

import "testing"

func TestTreeInsertAndGet(t *testing.T) {
	tr := New[string]()

	if _, updated := tr.Insert("/message", "first"); updated {
		t.Fatalf("expected first insert to not report an update")
	}
	if v, found := tr.Get("/message"); !found || v != "first" {
		t.Fatalf("expected to find /message=first, got %q (found=%v)", v, found)
	}
	if old, updated := tr.Insert("/message", "second"); !updated || old != "first" {
		t.Fatalf("expected second insert to report an update of %q, got %q (updated=%v)", "first", old, updated)
	}
}

func TestTreeLongestPrefix(t *testing.T) {
	tr := New[string]()
	tr.Insert("/message", "root")
	tr.Insert("/message/text", "leaf")

	if key, v, ok := tr.LongestPrefix("/message/text/extra"); !ok || key != "/message/text" || v != "leaf" {
		t.Fatalf("expected longest prefix /message/text=leaf, got %q=%q (ok=%v)", key, v, ok)
	}
	if key, v, ok := tr.LongestPrefix("/message/other"); !ok || key != "/message" || v != "root" {
		t.Fatalf("expected longest prefix /message=root, got %q=%q (ok=%v)", key, v, ok)
	}
}

func TestTreeNoMatch(t *testing.T) {
	tr := New[int]()
	if _, _, ok := tr.LongestPrefix("/anything"); ok {
		t.Fatalf("expected no match on an empty tree")
	}
	if _, ok := tr.Get("/anything"); ok {
		t.Fatalf("expected no match on an empty tree")
	}
}
