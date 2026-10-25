package main

import (
	"bytes"
	"errors"
	"go/format"
	"slices"
	"testing"
)

func TestNames(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) string
		in   string
		want string
	}{
		{"exported", exported, "chat_id", "ChatID"},
		{"exported with several initialisms", exported, "thumbnail_url", "ThumbnailURL"},
		{"exported plural ids", exported, "message_ids", "MessageIDs"},
		{"exported single word", exported, "text", "Text"},
		{"unexported", unexported, "chat_id", "chatID"},
		{"unexported starting with an initialism", unexported, "url", "url"},
		{"unexported keyword", unexported, "type", "typ"},
		{"unexported reserved name", unexported, "offset", "offset"},
		{"unexported clashing with method bodies", unexported, "ctx", "ctxArg"},
		{"type name", typeName, "MessageId", "MessageID"},
		{"type name with initialism inside", typeName, "RichTextUrl", "RichTextURL"},
		{"type name with digits", typeName, "InlineQueryResultMpeg4Gif", "InlineQueryResultMpeg4Gif"},
		{"type name with API", typeName, "TransactionPartnerTelegramApi", "TransactionPartnerTelegramAPI"},
		{"method name", methodName, "sendMessage", "SendMessage"},
		{"method name with initialism", methodName, "getWebhookInfo", "GetWebhookInfo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(tt.in); got != tt.want {
				t.Errorf("%s(%q) = %q, want %q", tt.name, tt.in, got, tt.want)
			}
		})
	}
}

func TestIs64(t *testing.T) {
	tests := []struct {
		field apiField
		want  bool
	}{
		{apiField{Name: "user_id"}, true},
		{apiField{Name: "sender_chat_id"}, true},
		{apiField{Name: "user_ids"}, true},
		{apiField{Name: "date"}, true},
		{apiField{Name: "until_date"}, true},
		{apiField{Name: "id", Description: "This number has at most 52 significant bits"}, true},
		{apiField{Name: "message_id"}, false},
		{apiField{Name: "offset"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.field.Name, func(t *testing.T) {
			if got := is64(tt.field); got != tt.want {
				t.Errorf("is64(%+v) = %v, want %v", tt.field, got, tt.want)
			}
		})
	}
}

func TestFindTag(t *testing.T) {
	str := goType{Expr: "string"}
	tests := []struct {
		name               string
		fields             []structField
		wantField, wantVal string
	}{
		{"always in quotes", []structField{{JSON: "status", Type: str, Required: true, Doc: `The member's status in the chat, always "creator"`}}, "status", "creator"},
		{"must be", []structField{{JSON: "type", Type: str, Required: true, Doc: "Type of the result, must be article"}}, "type", "article"},
		{"typographic quotes", []structField{{JSON: "type", Type: str, Required: true, Doc: "Type of the origin, always “user”"}}, "type", "user"},
		{"optional fields are not tags", []structField{{JSON: "type", Type: str, Doc: `always "x"`}}, "", ""},
		{"no tag", []structField{{JSON: "text", Type: str, Required: true, Doc: "Text of the message"}}, "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, val := findTag(&structDef{Fields: tt.fields})
			if field != tt.wantField || val != tt.wantVal {
				t.Errorf("findTag = %q, %q; want %q, %q", field, val, tt.wantField, tt.wantVal)
			}
		})
	}
}

func TestTextPairsAndTargets(t *testing.T) {
	f := func(names ...string) []apiField {
		var fields []apiField
		for _, n := range names {
			fields = append(fields, apiField{Name: n})
		}
		return fields
	}
	tests := []struct {
		name       string
		fields     []apiField
		wantPairs  map[string]string
		wantTarget bool
	}{
		{"text", f("chat_id", "text", "parse_mode"), map[string]string{"text": "parse_mode"}, false},
		{"caption", f("chat_id", "photo", "caption", "parse_mode"), map[string]string{"caption": "parse_mode"}, false},
		{"prefixed parse modes", f("question", "question_parse_mode", "explanation", "explanation_parse_mode"),
			map[string]string{"question": "question_parse_mode", "explanation": "explanation_parse_mode"}, false},
		{"edit target", f("chat_id", "message_id", "inline_message_id", "text", "parse_mode"), map[string]string{"text": "parse_mode"}, true},
		{"no pairs", f("chat_id", "user_id"), map[string]string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := textPairs(tt.fields)
			if len(got) != len(tt.wantPairs) {
				t.Errorf("textPairs = %v, want %v", got, tt.wantPairs)
			}
			for k, v := range tt.wantPairs {
				if got[k] != v {
					t.Errorf("textPairs[%s] = %q, want %q", k, got[k], v)
				}
			}
			if got := hasTarget(tt.fields); got != tt.wantTarget {
				t.Errorf("hasTarget = %v, want %v", got, tt.wantTarget)
			}
		})
	}
}

// TestSpec runs the analysis on the vendored spec: it must resolve every
// type and name without clashes.
func TestSpec(t *testing.T) {
	s, err := readSpec("../../apispec/api.html")
	if err != nil {
		t.Fatal(err)
	}
	g := newGenerator(s)
	if err := g.analyze(); err != nil {
		t.Fatalf("analyze: %v", err)
	}
	tests := []struct {
		name string
		ok   bool
	}{
		{"every method is generated", len(g.methods) == len(s.Methods)},
		{"ChatMember can be decoded", g.unions["ChatMember"] != nil && g.unions["ChatMember"].Decodable},
		{"InlineQueryResult is encode-only (its kinds share tags)", g.unions["InlineQueryResult"] != nil && !g.unions["InlineQueryResult"].Decodable},
		{"sendMessage takes the text as TextData", slices.ContainsFunc(findMethod(g, "sendMessage").Args, func(a arg) bool { return a.Name == "text" && a.Type == "TextData" })},
		{"editMessageText takes a MessageTarget", findMethod(g, "editMessageText").HasTarget},
		{"WithReplyMarkup takes any markup", g.options["ReplyMarkup"] != nil && g.options["ReplyMarkup"].ArgType == "ReplyMarkup"},
		{"WithInlineKeyboard works for edits", g.options["InlineKeyboard"] != nil && g.options["InlineKeyboard"].Applies["EditMessageText"] != ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.ok {
				t.Error("failed")
			}
		})
	}
}

func findMethod(g *generator, name string) *method {
	for _, m := range g.methods {
		if m.API.Name == name {
			return m
		}
	}
	return &method{}
}

// TestWriters runs every writer on the vendored spec: each must produce
// valid Go without reporting errors.
func TestWriters(t *testing.T) {
	s, err := readSpec("../../apispec/api.html")
	if err != nil {
		t.Fatal(err)
	}
	g := newGenerator(s)
	if err := g.analyze(); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name  string
		write func(*bytes.Buffer)
	}{
		{"types", g.writeTypes},
		{"unions", g.writeUnions},
		{"methods", g.writeMethods},
		{"options", g.writeOptions},
		{"method tests", g.writeMethodTests},
		{"shortcuts", g.writeShortcuts},
		{"shortcut tests", g.writeShortcutTests}, // after shortcuts
		{"sandbox results", g.writeDefaultResults},
		{"sandbox tests", g.writeSandboxTests},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.write(&buf)
			if _, err := format.Source(buf.Bytes()); err != nil {
				t.Errorf("invalid Go: %v", err)
			}
			if err := errors.Join(g.errs...); err != nil {
				t.Errorf("errors: %v", err)
			}
		})
	}
}
