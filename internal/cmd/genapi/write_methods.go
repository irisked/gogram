package main

import (
	"bytes"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func (m *method) signature(qualify func(string) string) (params, results string) {
	var ps []string
	for _, a := range m.Args {
		ps = append(ps, a.Name+" "+qualify(a.Type))
	}
	if len(m.Options) > 0 {
		ps = append(ps, "opts ..."+qualify(m.Name+"Option"))
	}
	params = strings.Join(ps, ", ")
	switch m.Return {
	case returnError:
		results = "error"
	case returnEdit:
		results = "(" + qualify("*Message") + ", error)"
	default:
		results = "(" + qualify(m.Result.Expr) + ", error)"
	}
	return params, results
}

func identity(s string) string { return s }

// zero returns the zero value literal of a Go type expression.
func zero(t string) string {
	switch {
	case t == "int" || t == "int64" || t == "float64":
		return "0"
	case t == "string":
		return `""`
	case t == "bool":
		return "false"
	case strings.HasPrefix(t, "*"), strings.HasPrefix(t, "[]"):
		return "nil"
	case t == "TextData" || t == "MessageTarget":
		return t + "{}"
	default:
		return "nil" // interfaces
	}
}

func (g *generator) writeMethods(buf *bytes.Buffer) {
	buf.WriteString("package telegram\n\nimport (\n\"context\"\n\"encoding/json\"\n)\n\n")
	buf.WriteString("var _ json.RawMessage // used by methods returning interfaces\n\n")
	for _, m := range g.methods {
		// Request type.
		comment(buf, "", fmt.Sprintf("%s is the request of [Client.%s]. Create it with [New%s] and send it with [Client.Do].", m.Name, m.Name, m.Name))
		fmt.Fprintf(buf, "type %s struct {\n", m.Name)
		for _, f := range m.Fields {
			if f.Name == "" {
				buf.WriteString("\t// The message to act on.\n\tMessageTarget\n")
				continue
			}
			if f.Doc != "" {
				comment(buf, "\t", f.Doc)
			}
			fmt.Fprintf(buf, "\t%s %s %s\n", f.Name, f.Type.Expr, f.tag())
		}
		buf.WriteString("}\n\n")
		fmt.Fprintf(buf, "// Method returns %q.\nfunc (*%s) Method() string { return %q }\n\n", m.API.Name, m.Name, m.API.Name)

		// Constructor.
		var params, assigns []string
		for _, a := range m.Args {
			params = append(params, a.Name+" "+a.Type)
			assigns = append(assigns, a.Assign)
		}
		fmt.Fprintf(buf, "// New%s creates a %s request with its required parameters.\n", m.Name, m.API.Name)
		fmt.Fprintf(buf, "func New%s(%s) *%s {\n\treturn &%s{%s}\n}\n\n", m.Name, strings.Join(params, ", "), m.Name, m.Name, strings.Join(assigns, ", "))

		// Option interface.
		if len(m.Options) > 0 {
			var names []string
			for _, o := range m.Options {
				names = append(names, "[With"+o+"]")
			}
			comment(buf, "", fmt.Sprintf("%sOption is an optional parameter of [Client.%s]: %s.", m.Name, m.Name, strings.Join(names, ", ")))
			fmt.Fprintf(buf, "type %sOption interface {\n\tapply%s(*%s)\n}\n\n", m.Name, m.Name, m.Name)
		}

		// Client method.
		sigParams, results := m.signature(identity)
		doc := append([]string(nil), m.Doc...)
		doc = append(doc, "", "See "+m.API.Href+".")
		comment(buf, "", doc...)
		fmt.Fprintf(buf, "func (c *Client) %s(ctx context.Context", m.Name)
		if sigParams != "" {
			buf.WriteString(", " + sigParams)
		}
		fmt.Fprintf(buf, ") %s {\n", results)
		var args []string
		for _, a := range m.Args {
			args = append(args, a.Name)
		}
		fmt.Fprintf(buf, "\treq := New%s(%s)\n", m.Name, strings.Join(args, ", "))
		if len(m.Options) > 0 {
			fmt.Fprintf(buf, "\tfor _, o := range opts {\n\t\tif o != nil {\n\t\t\to.apply%s(req)\n\t\t}\n\t}\n", m.Name)
		}
		g.writeReturn(buf, m)
		buf.WriteString("}\n\n")
	}
}

func (g *generator) writeReturn(buf *bytes.Buffer, m *method) {
	switch m.Return {
	case returnError:
		buf.WriteString("\treturn c.Do(ctx, req, nil)\n")
		return
	case returnEdit:
		buf.WriteString("\treturn c.processEdit(ctx, req)\n")
		return
	}
	r := m.Result
	switch {
	case r.isUnion() && r.Depth == 0:
		fmt.Fprintf(buf, "\tvar raw json.RawMessage\n\tif err := c.Do(ctx, req, &raw); err != nil {\n\t\treturn nil, err\n\t}\n\treturn decodeResult(req, raw, decode%s)\n", r.Union)
	case r.isUnion() && r.Depth == 1:
		fmt.Fprintf(buf, "\tvar raw []json.RawMessage\n\tif err := c.Do(ctx, req, &raw); err != nil {\n\t\treturn nil, err\n\t}\n\treturn decodeResults(req, raw, decode%s)\n", r.Union)
	case strings.HasPrefix(r.Expr, "*"):
		fmt.Fprintf(buf, "\tvar res %s\n\tif err := c.Do(ctx, req, &res); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &res, nil\n", r.Expr[1:])
	default:
		fmt.Fprintf(buf, "\tvar res %s\n\tif err := c.Do(ctx, req, &res); err != nil {\n\t\treturn %s, err\n\t}\n\treturn res, nil\n", r.Expr, zero(r.Expr))
	}
}

func (g *generator) writeOptions(buf *bytes.Buffer) {
	buf.WriteString("package telegram\n\n")
	for _, name := range sortedKeys(g.options) {
		o := g.options[name]
		reqs := sortedKeys(o.Applies)
		var accepted []string
		for _, r := range reqs {
			accepted = append(accepted, "[Client."+r+"]")
		}
		comment(buf, "", fmt.Sprintf("%sOption is the option created by [With%s]. It is accepted by %s.", name, name, strings.Join(accepted, ", ")))
		switch {
		case o.ArgType == "":
			fmt.Fprintf(buf, "type %sOption struct{}\n\n", name)
		case o.Variadic:
			fmt.Fprintf(buf, "type %sOption struct {\n\tv []%s\n}\n\n", name, o.ArgType)
		default:
			fmt.Fprintf(buf, "type %sOption struct {\n\tv %s\n}\n\n", name, o.ArgType)
		}

		doc := fmt.Sprintf("With%s sets %s: %s", name, o.JSON, o.Doc)
		if slices.ContainsFunc(extraOptions, func(e extraOption) bool { return e.Name == name }) {
			doc = fmt.Sprintf("With%s %s", name, o.Doc)
		}
		comment(buf, "", doc)
		switch {
		case o.ArgType == "":
			fmt.Fprintf(buf, "func With%s() %sOption { return %sOption{} }\n\n", name, name, name)
		case o.Variadic:
			fmt.Fprintf(buf, "func With%s(%s ...%s) %sOption { return %sOption{v: %s} }\n\n", name, o.ArgName, o.ArgType, name, name, o.ArgName)
		default:
			fmt.Fprintf(buf, "func With%s(%s %s) %sOption { return %sOption{v: %s} }\n\n", name, o.ArgName, o.ArgType, name, name, o.ArgName)
		}
		for _, r := range reqs {
			fmt.Fprintf(buf, "func (o %sOption) apply%s(r *%s) { %s }\n", name, r, r, o.Applies[r])
		}
		buf.WriteString("\n")
	}
}

func (g *generator) writeMethodTests(buf *bytes.Buffer) {
	buf.WriteString(`package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGeneratedRequests checks every request built with its constructor:
// its method name, and that it encodes to JSON with all required
// parameters present.
func TestGeneratedRequests(t *testing.T) {
	tests := []struct {
		method   string
		req      Request
		required []string
	}{
`)
	for _, m := range g.methods {
		var args, required []string
		for _, a := range m.Args {
			args = append(args, zero(a.Type))
		}
		for _, f := range m.Fields {
			if f.Required && f.Name != "" {
				required = append(required, fmt.Sprintf("%q", f.JSON))
			}
		}
		fmt.Fprintf(buf, "\t\t{%q, New%s(%s), []string{%s}},\n", m.API.Name, m.Name, strings.Join(args, ", "), strings.Join(required, ", "))
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			if got := tt.req.Method(); got != tt.method {
				t.Errorf("Method() = %q, want %q", got, tt.method)
			}
			body, contentType, err := encodeRequest(tt.req)
			if err != nil {
				t.Fatalf("encode: %v", err)
			}
			if contentType != "application/json" {
				t.Fatalf("Content-Type = %q, want application/json", contentType)
			}
			var fields map[string]json.RawMessage
			if err := json.NewDecoder(body).Decode(&fields); err != nil {
				t.Fatalf("decode: %v", err)
			}
			for _, name := range tt.required {
				if _, ok := fields[name]; !ok {
					t.Errorf("required parameter %s is missing from %v", name, fields)
				}
			}
		})
	}
}
`)
	g.writeClientTests(buf)
	g.writeUnionTests(buf)
}

var reExported = regexp.MustCompile(`\b([A-Z]\w*)`)

func qualifyTelegram(expr string) string {
	return reExported.ReplaceAllString(expr, "telegram.$1")
}

// writeShortcuts writes the gogram.Update methods that call a Client
// method for the update's chat or callback message.
func (g *generator) writeShortcuts(buf *bytes.Buffer) {
	buf.WriteString("package gogram\n\nimport (\n\"context\"\n\n\"github.com/irisked/gogram/telegram\"\n)\n\n")
	byName := map[string]*method{}
	for _, m := range g.methods {
		byName[m.API.Name] = m
	}
	firstIsChat := func(m *method) bool {
		return len(m.Args) > 0 && m.Args[0].Name == "chatID" && m.Args[0].Type == "int64"
	}

	buf.WriteString("// --- Sending to the update's chat ---\n\n")
	for _, m := range g.methods {
		if !strings.HasPrefix(m.API.Name, "send") || replyExclude[m.API.Name] || !firstIsChat(m) {
			continue
		}
		name := "ReplyWith" + strings.TrimPrefix(m.Name, "Send")
		if m.API.Name == "sendMessage" {
			name = "Reply"
		}
		g.writeShortcut(buf, m, name, "chatID", "u.chatID()",
			fmt.Sprintf("%s calls [telegram.Client.%s] for the update's chat.", name, m.Name))
	}

	buf.WriteString("// --- Acting on the update's chat ---\n\n")
	for _, api := range chatShortcuts {
		m := byName[api]
		if m == nil || !firstIsChat(m) {
			g.errorf("chat shortcut %s: no such method taking chat_id first", api)
			continue
		}
		g.writeShortcut(buf, m, m.Name, "chatID", "u.chatID()",
			fmt.Sprintf("%s calls [telegram.Client.%s] for the update's chat.", m.Name, m.Name))
	}

	buf.WriteString("// --- Editing the message whose button was pressed ---\n\n")
	for _, m := range g.methods {
		if !m.HasTarget || !(strings.HasPrefix(m.API.Name, "editMessage") || m.API.Name == "stopMessageLiveLocation") {
			continue
		}
		g.writeShortcut(buf, m, m.Name, "target", "u.callbackTarget()",
			fmt.Sprintf("%s calls [telegram.Client.%s] on the message whose button was pressed. It returns the edited message, or nil for inline messages.", m.Name, m.Name))
	}
}

func (g *generator) writeShortcut(buf *bytes.Buffer, m *method, name, first, source, doc string) {
	g.shortcuts = append(g.shortcuts, shortcut{name: name, method: m})
	rest := &method{Args: m.Args[1:], Options: m.Options, Return: m.Return, Result: m.Result, Name: m.Name}
	params, results := rest.signature(qualifyTelegram)
	comment(buf, "", doc)
	fmt.Fprintf(buf, "func (u *Update) %s(ctx context.Context", name)
	if params != "" {
		buf.WriteString(", " + params)
	}
	fmt.Fprintf(buf, ") %s {\n", results)
	fail := "return err"
	if m.Return != returnError {
		r := m.Result.Expr
		if m.Return == returnEdit {
			r = "*Message"
		}
		fail = "return " + zero(r) + ", err"
	}
	fmt.Fprintf(buf, "\t%s, err := %s\n\tif err != nil {\n\t\t%s\n\t}\n", first, source, fail)
	args := []string{"ctx", first}
	for _, a := range rest.Args {
		args = append(args, a.Name)
	}
	if len(m.Options) > 0 {
		args = append(args, "opts...")
	}
	fmt.Fprintf(buf, "\treturn u.client.%s(%s)\n}\n\n", m.Name, strings.Join(args, ", "))
}

// resultJSON returns a plausible result of a method, for tests and for
// the default answers of sandbox.
func (g *generator) resultJSON(m *method) string {
	const message = `{"message_id":1,"date":1,"chat":{"id":1,"type":"private"}}`
	switch m.Return {
	case returnError:
		return "true"
	case returnEdit:
		return message
	}
	r := m.Result
	if r.isUnion() {
		u := g.unions[r.Union]
		obj := "{}"
		if u.Decodable {
			obj = fmt.Sprintf(`{%q:%q}`, u.TagField, u.Tags[u.Members[0]])
		}
		if r.Depth == 1 {
			return "[" + obj + "]"
		}
		return obj
	}
	switch {
	case r.Expr == "*Message":
		return message
	case strings.HasPrefix(r.Expr, "[]"):
		return "[]"
	case strings.HasPrefix(r.Expr, "*"):
		return "{}"
	case r.Expr == "string":
		return `""`
	case r.Expr == "bool":
		return "true"
	default:
		return "0"
	}
}

// sample returns a non-zero value literal of a Go type, for tests.
func (g *generator) sample(t string) string {
	switch t {
	case "int", "int64":
		return "1"
	case "float64":
		return "1.5"
	case "string":
		return `"x"`
	case "bool":
		return "true"
	case "TextData":
		return `TextData{Text: "x"}`
	case "*InputFile":
		return `FileID("x")`
	case "ReplyMarkup":
		return "&InlineKeyboardMarkup{}"
	case "UpdateType":
		return "UpdateMessage"
	case "ChatID":
		return "ChatID{ID: 1}"
	}
	if u, ok := g.unions[t]; ok {
		return "&" + u.Members[0] + "{}"
	}
	switch {
	case strings.HasPrefix(t, "*") && strings.ToLower(t[1:2]) == t[1:2]:
		return "new(" + t[1:] + ")"
	case strings.HasPrefix(t, "*"):
		return "&" + t[1:] + "{}"
	case strings.HasPrefix(t, "[]"):
		return t + "{" + g.sample(t[2:]) + "}"
	default:
		return t + "{}"
	}
}

// writeGeneratedTests appends the tests of client methods and options.
func (g *generator) writeClientTests(buf *bytes.Buffer) {
	buf.WriteString(`
// TestGeneratedClientMethods calls every Client method against a server
// answering with a result of the right type.
func TestGeneratedClientMethods(t *testing.T) {
	results := map[string]string{
`)
	for _, m := range g.methods {
		fmt.Fprintf(buf, "\t\t%q: %q,\n", m.API.Name, g.resultJSON(m))
	}
	buf.WriteString(`	}
	var called string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		io.WriteString(w, ` + "`" + `{"ok":true,"result":` + "`" + `+results[called]+"}")
	}))
	defer srv.Close()
	c := New("token", WithBaseURL(srv.URL))
	ctx := context.Background()

	tests := []struct {
		method string
		call   func() error
	}{
`)
	for _, m := range g.methods {
		var args []string
		for _, a := range m.Args {
			args = append(args, zero(a.Type))
		}
		call := fmt.Sprintf("c.%s(ctx", m.Name)
		if len(args) > 0 {
			call += ", " + strings.Join(args, ", ")
		}
		call += ")"
		if m.Return == returnError {
			fmt.Fprintf(buf, "\t\t{%q, func() error { return %s }},\n", m.API.Name, call)
		} else {
			fmt.Fprintf(buf, "\t\t{%q, func() error { _, err := %s; return err }},\n", m.API.Name, call)
		}
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			if err := tt.call(); err != nil {
				t.Fatalf("error = %v", err)
			}
			if called != tt.method {
				t.Errorf("called %s, want %s", called, tt.method)
			}
		})
	}
}

// TestGeneratedOptions applies every option to every request accepting
// it and checks that it sets its parameter.
func TestGeneratedOptions(t *testing.T) {
	tests := []struct {
		name  string
		param string
		req   func() Request
	}{
`)
	byName := map[string]*method{}
	for _, m := range g.methods {
		byName[m.Name] = m
	}
	for _, name := range sortedKeys(g.options) {
		o := g.options[name]
		var arg string
		switch {
		case o.ArgType == "":
		default:
			arg = g.sample(o.ArgType)
		}
		for _, r := range sortedKeys(o.Applies) {
			m := byName[r]
			var args []string
			for _, a := range m.Args {
				args = append(args, zero(a.Type))
			}
			fmt.Fprintf(buf, "\t\t{%q, %q, func() Request { r := New%s(%s); With%s(%s).apply%s(r); return r }},\n",
				"With"+name+"/"+m.API.Name, o.JSON, r, strings.Join(args, ", "), name, arg, r)
		}
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _, err := encodeRequest(tt.req())
			if err != nil {
				t.Fatalf("encode: %v", err)
			}
			var fields map[string]json.RawMessage
			if err := json.NewDecoder(body).Decode(&fields); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if _, ok := fields[tt.param]; !ok {
				t.Errorf("parameter %s is not set: %v", tt.param, fields)
			}
		})
	}
}
`)
}

// writeShortcutTests writes the test of the generated gogram.Update
// shortcuts.
func (g *generator) writeShortcutTests(buf *bytes.Buffer) {
	buf.WriteString(`package gogram_test

import (
	"context"
	"testing"

	"github.com/irisked/gogram"
	"github.com/irisked/gogram/sandbox"
	"github.com/irisked/gogram/telegram"
)

var _ telegram.TextData // used by some calls

// TestGeneratedShortcuts calls every generated Update shortcut on a
// callback query, which has both a chat and a message.
func TestGeneratedShortcuts(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(ctx context.Context, u *gogram.Update) error
	}{
`)
	for _, s := range g.shortcuts {
		var args []string
		for _, a := range s.method.Args[1:] {
			args = append(args, qualifyTelegram(zero(a.Type)))
		}
		call := fmt.Sprintf("u.%s(ctx", s.name)
		if len(args) > 0 {
			call += ", " + strings.Join(args, ", ")
		}
		call += ")"
		if s.method.Return == returnError {
			fmt.Fprintf(buf, "\t\t{%q, %q, func(ctx context.Context, u *gogram.Update) error { return %s }},\n", s.name, s.method.API.Name, call)
		} else {
			fmt.Fprintf(buf, "\t\t{%q, %q, func(ctx context.Context, u *gogram.Update) error { _, err := %s; return err }},\n", s.name, s.method.API.Name, call)
		}
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			u := gogram.NewUpdate(srv.Client(), sandbox.Callback("x"))
			if err := tt.call(t.Context(), u); err != nil {
				t.Fatalf("error = %v", err)
			}
			if reqs := srv.Requests(); len(reqs) != 1 || reqs[0].Method != tt.method {
				t.Errorf("requests = %v, want one %s", reqs, tt.method)
			}
		})
	}
}
`)
}

// writeDefaultResults writes the default answers of sandbox.Server.
func (g *generator) writeDefaultResults(buf *bytes.Buffer) {
	buf.WriteString("package sandbox\n\n// defaultResults are plausible results of every method, by name.\nvar defaultResults = map[string]string{\n")
	for _, m := range g.methods {
		fmt.Fprintf(buf, "\t%q: %q,\n", m.API.Name, g.resultJSON(m))
	}
	buf.WriteString("}\n")
}

// writeUnionTests appends the tests of union decoding and encoding.
func (g *generator) writeUnionTests(buf *bytes.Buffer) {
	buf.WriteString(`
// TestGeneratedUnions decodes every kind of every interface type from
// its discriminator, and encodes it back.
func TestGeneratedUnions(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		decode func([]byte) (any, error)
		want   any
	}{
`)
	for _, name := range sortedKeys(g.unions) {
		u := g.unions[name]
		if !u.Decodable {
			continue
		}
		for _, m := range u.Members {
			fmt.Fprintf(buf, "\t\t{%q, %q, func(b []byte) (any, error) { return decode%s(b) }, &%s{}},\n",
				name+"/"+u.Tags[m], fmt.Sprintf(`{%q:%q}`, u.TagField, u.Tags[m]), name, m)
		}
		fmt.Fprintf(buf, "\t\t{%q, %q, func(b []byte) (any, error) { return decode%s(b) }, &UnknownVariant{Raw: json.RawMessage(%q)}},\n",
			name+"/unknown", fmt.Sprintf(`{%q:"future_kind"}`, u.TagField), name, fmt.Sprintf(`{%q:"future_kind"}`, u.TagField))
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.decode([]byte(tt.json))
			if err != nil {
				t.Fatalf("decode: %v", err)
			}
			if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.want) {
				t.Fatalf("decoded %T, want %T", got, tt.want)
			}
			b, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("encode: %v", err)
			}
			// The discriminator comes first; required fields follow.
			if !strings.HasPrefix(string(b), strings.TrimSuffix(tt.json, "}")) {
				t.Errorf("encoded %s, want it to start with %s", b, strings.TrimSuffix(tt.json, "}"))
			}
		})
	}
}

// TestGeneratedUnionFields decodes every interface-typed field of every
// type, and encodes it back.
func TestGeneratedUnionFields(t *testing.T) {
	tests := []struct {
		name string
		json string
		tag  string // must survive decoding and encoding again
		new  func() any
	}{
`)
	for _, name := range sortedKeys(g.structs) {
		s := g.structs[name]
		for _, f := range s.Fields {
			if !f.Type.isUnion() || !g.unions[f.Type.Union].Decodable {
				continue
			}
			u := g.unions[f.Type.Union]
			v := fmt.Sprintf(`{%q:%q}`, u.TagField, u.Tags[u.Members[0]])
			for range f.Type.Depth {
				v = "[" + v + "]"
			}
			tag := fmt.Sprintf(`%q:%q`, u.TagField, u.Tags[u.Members[0]])
			fmt.Fprintf(buf, "\t\t{%q, %q, %q, func() any { return new(%s) }},\n", name+"."+f.JSON, fmt.Sprintf(`{%q:%s}`, f.JSON, v), tag, name)
		}
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.new()
			if err := json.Unmarshal([]byte(tt.json), v); err != nil {
				t.Fatalf("decode: %v", err)
			}
			b, err := json.Marshal(v)
			if err != nil {
				t.Fatalf("encode: %v", err)
			}
			if !strings.Contains(string(b), tt.tag) {
				t.Errorf("encoded %s, want it to contain %s", b, tt.tag)
			}
		})
	}
}
`)
}

// writeSandboxTests writes the test of the sandbox server's default
// answers: every method must get one its Client method can decode.
func (g *generator) writeSandboxTests(buf *bytes.Buffer) {
	buf.WriteString(`package sandbox_test

import (
	"context"
	"testing"

	"github.com/irisked/gogram/sandbox"
	"github.com/irisked/gogram/telegram"
)

var _ telegram.TextData // used by some calls

// TestGeneratedDefaultAnswers calls every Client method against the
// server: each must succeed with the default answer, and be recorded.
func TestGeneratedDefaultAnswers(t *testing.T) {
	tests := []struct {
		method string
		call   func(ctx context.Context, c *telegram.Client) error
	}{
`)
	for _, m := range g.methods {
		if m.API.Name == "getUpdates" {
			continue // answered from the queue; see TestServerGetUpdates
		}
		var args []string
		for _, a := range m.Args {
			args = append(args, qualifyTelegram(zero(a.Type)))
		}
		call := fmt.Sprintf("c.%s(ctx", m.Name)
		if len(args) > 0 {
			call += ", " + strings.Join(args, ", ")
		}
		call += ")"
		if m.Return == returnError {
			fmt.Fprintf(buf, "\t\t{%q, func(ctx context.Context, c *telegram.Client) error { return %s }},\n", m.API.Name, call)
		} else {
			fmt.Fprintf(buf, "\t\t{%q, func(ctx context.Context, c *telegram.Client) error { _, err := %s; return err }},\n", m.API.Name, call)
		}
	}
	buf.WriteString(`	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			srv := sandbox.NewServer(t)
			if err := tt.call(t.Context(), srv.Client()); err != nil {
				t.Fatalf("error = %v", err)
			}
			if calls := srv.Calls(tt.method); len(calls) != 1 {
				t.Errorf("recorded %d %s calls, want 1", len(calls), tt.method)
			}
		})
	}
}
`)
}
