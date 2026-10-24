package router

import "testing"

func TestTreeInsert(t *testing.T) {
	tests := []struct {
		name        string
		existing    map[string]string
		key, value  string
		wantOld     string
		wantUpdated bool
	}{
		{name: "into empty tree", key: "/message", value: "v"},
		{name: "new key", existing: map[string]string{"/message": "a"}, key: "/message/text", value: "b"},
		{name: "key sharing a prefix", existing: map[string]string{"/message/text": "a"}, key: "/message/photo", value: "b"},
		{name: "existing key", existing: map[string]string{"/message": "first"}, key: "/message", value: "second", wantOld: "first", wantUpdated: true},
		{name: "empty key", existing: map[string]string{"/message": "a"}, key: "", value: "root"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New[string]()
			for k, v := range tt.existing {
				tr.Insert(k, v)
			}

			old, updated := tr.Insert(tt.key, tt.value)
			if old != tt.wantOld || updated != tt.wantUpdated {
				t.Errorf("Insert(%q) = %q, %v; want %q, %v", tt.key, old, updated, tt.wantOld, tt.wantUpdated)
			}
			if got, ok := tr.Get(tt.key); !ok || got != tt.value {
				t.Errorf("Get(%q) after Insert = %q, %v; want %q, true", tt.key, got, ok, tt.value)
			}
		})
	}
}

func TestTreeGet(t *testing.T) {
	tr := New[string]()
	tr.Insert("/message", "message")
	tr.Insert("/message/text", "text")

	tests := []struct {
		key    string
		want   string
		wantOK bool
	}{
		{"/message", "message", true},
		{"/message/text", "text", true},
		{"/message/", "", false},
		{"/mess", "", false},
		{"/message/text/extra", "", false},
		{"", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			got, ok := tr.Get(tt.key)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("Get(%q) = %q, %v; want %q, %v", tt.key, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestTreeLongestPrefix(t *testing.T) {
	tests := []struct {
		name             string
		keys             []string
		path             string
		wantKey, wantVal string
		wantOK           bool
	}{
		{"empty tree", nil, "/anything", "", "", false},
		{"exact key", []string{"/message"}, "/message", "/message", "/message", true},
		{"deepest key wins", []string{"/message", "/message/text"}, "/message/text/extra", "/message/text", "/message/text", true},
		{"falls back to shorter key", []string{"/message", "/message/text"}, "/message/other", "/message", "/message", true},
		{"ignores segment boundaries", []string{"/start"}, "/startgame", "/start", "/start", true},
		{"no key is a prefix", []string{"/message"}, "/callback", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New[string]()
			for _, k := range tt.keys {
				tr.Insert(k, k)
			}

			key, v, ok := tr.LongestPrefix(tt.path)
			if key != tt.wantKey || v != tt.wantVal || ok != tt.wantOK {
				t.Errorf("LongestPrefix(%q) = %q, %q, %v; want %q, %q, %v",
					tt.path, key, v, ok, tt.wantKey, tt.wantVal, tt.wantOK)
			}
		})
	}
}

func TestTreeMatch(t *testing.T) {
	routes := []string{"", "/message", "/message/command/start"}

	tests := []struct {
		name             string
		routes           []string
		path             string
		wantKey, wantVal string
		wantOK           bool
	}{
		{"empty tree", nil, "/message", "", "", false},
		{"no route matches", []string{"/message"}, "/callback_query/x", "", "", false},
		{"exact route", routes, "/message/command/start", "/message/command/start", "/message/command/start", true},
		{"route followed by more segments", routes, "/message/command/start/extra", "/message/command/start", "/message/command/start", true},
		{"route as part of a segment", routes, "/message/command/startgame", "/message", "/message", true},
		{"parent route", routes, "/message/text", "/message", "/message", true},
		{"empty route catches the rest", routes, "/messages", "", "", true},
		{"empty route for another type", routes, "/callback_query/x", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New[string]()
			for _, r := range tt.routes {
				tr.Insert(r, r)
			}

			key, v, ok := tr.Match(tt.path)
			if key != tt.wantKey || v != tt.wantVal || ok != tt.wantOK {
				t.Errorf("Match(%q) = %q, %q, %v; want %q, %q, %v",
					tt.path, key, v, ok, tt.wantKey, tt.wantVal, tt.wantOK)
			}
		})
	}
}
