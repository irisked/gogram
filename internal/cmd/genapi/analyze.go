package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
)

// goType is a resolved Go type.
type goType struct {
	Expr  string // e.g. "*Message", "[]PhotoSize", "ChatMember"
	Union string // union interface involved in Expr, if any
	Depth int    // slice depth around Union
}

func (t goType) isUnion() bool { return t.Union != "" }

type structField struct {
	Name     string // Go name; "" for the embedded MessageTarget
	JSON     string
	Doc      string
	Type     goType
	Required bool
}

func (f structField) tag() string {
	switch {
	case f.Name == "":
		return ""
	case f.Required:
		return fmt.Sprintf("`json:%q`", f.JSON)
	case f.Type.Expr == "ChatID":
		return fmt.Sprintf("`json:%q`", f.JSON+",omitzero")
	default:
		return fmt.Sprintf("`json:%q`", f.JSON+",omitempty")
	}
}

type structDef struct {
	Name   string
	Doc    []string
	Href   string
	Fields []structField
	// Set for union members: the discriminator stripped from Fields.
	TagField, TagValue string
	Unions             []string
}

type union struct {
	Name      string
	Doc       []string
	Href      string
	Members   []string // Go type names
	TagField  string   // "" if members can't be told apart when decoding
	Tags      map[string]string
	RichText  bool // also a string or an array of RichText
	Decodable bool // members can be told apart by TagField
}

// arg is a positional parameter of a method.
type arg struct {
	Name   string
	Type   string
	Assign string // composite literal entry setting the request field
}

type method struct {
	API       *apiMethod
	Name      string // Go name: request type and Client method
	Doc       []string
	Fields    []structField
	Args      []arg
	Options   []string // option names, sorted
	Return    returnKind
	Result    goType // for returnValue
	HasTarget bool
}

type returnKind int

const (
	returnError returnKind = iota // True
	returnValue                   // an object, list or primitive
	returnEdit                    // Message, or True for inline messages
)

type option struct {
	Name     string // without "With"
	ArgName  string
	ArgType  string // "" for a flag
	Variadic bool
	Doc      string
	Applies  map[string]string // request type -> statement
	JSON     string
}

type generator struct {
	spec      *spec
	structs   map[string]*structDef // Go name -> struct
	unions    map[string]*union     // Go name -> union
	unionOf   map[string]string     // spec union name -> Go name
	methods   []*method
	options   map[string]*option
	shortcuts []shortcut // filled while writing them
	errs      []error
}

type shortcut struct {
	name   string
	method *method
}

func newGenerator(s *spec) *generator {
	return &generator{
		spec:    s,
		structs: map[string]*structDef{},
		unions:  map[string]*union{},
		unionOf: map[string]string{},
		options: map[string]*option{},
	}
}

func (g *generator) errorf(format string, args ...any) {
	g.errs = append(g.errs, fmt.Errorf(format, args...))
}

func sortedKeys[M ~map[string]V, V any](m M) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (g *generator) analyze() error {
	// Unions first: field types refer to them.
	for _, name := range sortedKeys(g.spec.Types) {
		t := g.spec.Types[name]
		if len(t.Subtypes) == 0 || skipTypes[name] {
			continue
		}
		g.addUnion(typeName(name), t.Description, t.Href, t.Subtypes)
		g.unionOf[name] = typeName(name)
	}
	for _, name := range sortedKeys(g.spec.Types) {
		t := g.spec.Types[name]
		if len(t.Subtypes) > 0 || skipTypes[name] {
			continue
		}
		g.addStruct(t)
	}
	g.assignTags()
	for _, name := range sortedKeys(g.spec.Methods) {
		g.addMethod(g.spec.Methods[name])
	}
	g.addExtraOptions()
	for _, m := range g.methods {
		sort.Strings(m.Options)
	}
	g.checkNames()
	return errors.Join(g.errs...)
}

func (g *generator) addUnion(name string, doc []string, href string, specMembers []string) *union {
	u := &union{Name: name, Doc: doc, Href: href, Tags: map[string]string{}}
	for _, m := range specMembers {
		switch {
		case m == "String" || strings.HasPrefix(m, "Array of "):
			// RichText: plain strings and arrays are rich texts too.
			u.RichText = true
		default:
			u.Members = append(u.Members, typeName(m))
		}
	}
	g.unions[name] = u
	return u
}

func (g *generator) addStruct(t *apiType) {
	name := typeName(t.Name)
	s := &structDef{Name: name, Doc: typeDoc(name, t.Description), Href: t.Href}
	for _, f := range t.Fields {
		s.Fields = append(s.Fields, structField{
			Name:     exported(f.Name),
			JSON:     f.Name,
			Doc:      docText(f.Description),
			Type:     g.resolve(t.Name, f),
			Required: f.Required,
		})
	}
	g.structs[name] = s
}

// assignTags finds the discriminator of every union member and strips it
// from the member's fields: MarshalJSON adds it back. A union can be
// decoded when all members share the discriminator field and have
// distinct values.
func (g *generator) assignTags() {
	for _, uname := range sortedKeys(g.unions) {
		u := g.unions[uname]
		fields := map[string]bool{}
		values := map[string]bool{}
		unique := true
		for _, m := range u.Members {
			s := g.structs[m]
			if s == nil {
				g.errorf("union %s: unknown member %s", uname, m)
				continue
			}
			s.Unions = append(s.Unions, uname)
			field, value := findTag(s)
			fields[field] = true
			if field == "" {
				continue
			}
			if values[value] {
				unique = false
			}
			values[value] = true
			u.Tags[m] = value
			if s.TagField != "" && (s.TagField != field || s.TagValue != value) {
				g.errorf("%s has discriminator %s=%s in one union and %s=%s in another", m, s.TagField, s.TagValue, field, value)
			}
			s.TagField, s.TagValue = field, value
		}
		if len(fields) == 1 && !fields[""] {
			for f := range fields {
				u.TagField = f
			}
			u.Decodable = unique
		}
	}
	for _, s := range g.structs {
		if s.TagField == "" {
			continue
		}
		s.Fields = slices.DeleteFunc(s.Fields, func(f structField) bool { return f.JSON == s.TagField })
	}
}

func findTag(s *structDef) (field, value string) {
	for _, f := range s.Fields {
		if f.Type.Expr != "string" || !f.Required {
			continue
		}
		if m := reTag.FindStringSubmatch(f.Doc); m != nil {
			return f.JSON, m[1]
		}
	}
	return "", ""
}

// resolve maps the spec types of a field to a Go type.
func (g *generator) resolve(owner string, f apiField) goType {
	if t, ok := fieldTypes[owner+"."+f.Name]; ok {
		return goType{Expr: t}
	}
	if t, ok := fieldTypes["*."+f.Name]; ok {
		return goType{Expr: t}
	}
	set := strings.Join(f.Types, " | ")
	switch {
	case set == "String" && strings.Contains(f.Description, "attach://"):
		// Files inside objects: a file_id, a URL or an upload.
		return goType{Expr: "*InputFile"}
	case len(f.Types) == 1:
		return g.single(f.Types[0], f, false)
	case set == "Integer | String":
		return goType{Expr: "ChatID"}
	case set == "InputFile | String":
		return goType{Expr: "*InputFile"}
	case set == replyMarkupTypes:
		return goType{Expr: "ReplyMarkup"}
	}
	allArrays := true
	var elems []string
	for _, t := range f.Types {
		e, ok := strings.CutPrefix(t, "Array of ")
		allArrays = allArrays && ok
		elems = append(elems, e)
	}
	if allArrays {
		u := g.commonUnion(owner, f.Name, elems)
		return goType{Expr: "[]" + u, Union: u, Depth: 1}
	}
	u := g.commonUnion(owner, f.Name, f.Types)
	return goType{Expr: u, Union: u}
}

func (g *generator) single(t string, f apiField, elem bool) goType {
	switch t {
	case "Integer":
		if is64(f) {
			return goType{Expr: "int64"}
		}
		return goType{Expr: "int"}
	case "Float":
		return goType{Expr: "float64"}
	case "String":
		return goType{Expr: "string"}
	case "Boolean", "True":
		return goType{Expr: "bool"}
	case "InputFile":
		return goType{Expr: "*InputFile"}
	case "MaybeInaccessibleMessage":
		if elem {
			return goType{Expr: "Message"}
		}
		return goType{Expr: "*Message"}
	}
	if e, ok := strings.CutPrefix(t, "Array of "); ok {
		inner := g.single(e, f, true)
		inner.Expr = "[]" + inner.Expr
		if inner.isUnion() {
			inner.Depth++
		}
		return inner
	}
	if u, ok := g.unionOf[t]; ok {
		return goType{Expr: u, Union: u}
	}
	if _, ok := g.spec.Types[t]; !ok {
		g.errorf("field %s: unknown type %q", f.Name, t)
	}
	if elem {
		return goType{Expr: typeName(t)}
	}
	return goType{Expr: "*" + typeName(t)}
}

// commonUnion returns the smallest union containing all the given spec
// types, or declares a new one.
func (g *generator) commonUnion(owner, field string, types []string) string {
	want := make([]string, len(types))
	for i, t := range types {
		want[i] = typeName(t)
	}
	best := ""
	for _, name := range sortedKeys(g.unions) {
		u := g.unions[name]
		if !containsAll(u.Members, want) {
			continue
		}
		if best == "" || len(u.Members) < len(g.unions[best].Members) {
			best = name
		}
	}
	if best != "" {
		return best
	}
	name, ok := unionNames[owner+"."+field]
	if !ok {
		g.errorf("%s.%s: no union holds %v; name one in unionNames", owner, field, types)
		return "any"
	}
	if _, exists := g.unions[name]; !exists {
		g.addUnion(name, []string{fmt.Sprintf("%s is the type of %s.%s: one of the types listed below.", name, typeName(owner), exported(field))}, "", types)
	}
	return name
}

func containsAll(have, want []string) bool {
	for _, w := range want {
		if !slices.Contains(have, w) {
			return false
		}
	}
	return true
}

// textPairs maps each text field of a method to its parse mode field.
func textPairs(fields []apiField) map[string]string {
	names := map[string]bool{}
	for _, f := range fields {
		names[f.Name] = true
	}
	pairs := map[string]string{}
	for _, f := range fields {
		switch {
		case f.Name == "parse_mode" && names["text"]:
			pairs["text"] = f.Name
		case f.Name == "parse_mode" && names["caption"]:
			pairs["caption"] = f.Name
		case strings.HasSuffix(f.Name, "_parse_mode"):
			if text := strings.TrimSuffix(f.Name, "_parse_mode"); names[text] {
				pairs[text] = f.Name
			}
		}
	}
	return pairs
}

var targetFields = []string{"chat_id", "message_id", "inline_message_id"}

// hasTarget reports whether a method identifies a message by either
// chat_id and message_id or inline_message_id: those become a
// MessageTarget.
func hasTarget(fields []apiField) bool {
	n := 0
	for _, f := range fields {
		if slices.Contains(targetFields, f.Name) && !f.Required {
			n++
		}
	}
	return n == len(targetFields)
}

func (g *generator) addMethod(am *apiMethod) {
	m := &method{API: am, Name: methodName(am.Name), HasTarget: hasTarget(am.Fields)}
	m.Doc = append([]string{fmt.Sprintf("%s calls %s.", m.Name, am.Name)}, am.Description...)
	pairs := textPairs(am.Fields)
	parseModes := map[string]bool{}
	for _, pm := range pairs {
		parseModes[pm] = true
	}
	if m.HasTarget {
		m.Args = append(m.Args, arg{Name: "target", Type: "MessageTarget", Assign: "MessageTarget: target"})
		m.Fields = append(m.Fields, structField{Type: goType{Expr: "MessageTarget"}})
	}
	for _, f := range am.Fields {
		if m.HasTarget && slices.Contains(targetFields, f.Name) {
			continue
		}
		sf := structField{
			Name:     exported(f.Name),
			JSON:     f.Name,
			Doc:      docText(f.Description),
			Type:     g.resolve(am.Name, f),
			Required: f.Required,
		}
		if sf.Name == "Method" {
			g.errorf("%s: parameter %s clashes with the Method method", am.Name, f.Name)
		}
		m.Fields = append(m.Fields, sf)
		if parseModes[f.Name] {
			continue // set together with its text
		}
		_, isText := pairs[f.Name]
		pm := exported(pairs[f.Name])
		if f.Required || positional[am.Name+"."+f.Name] {
			a := arg{Name: unexported(f.Name), Type: sf.Type.Expr}
			switch {
			case isText:
				a.Type = "TextData"
				a.Assign = fmt.Sprintf("%s: %s.Text, %s: %s.ParseMode", sf.Name, a.Name, pm, a.Name)
			case sf.Type.Expr == "ChatID":
				a.Type = "int64"
				a.Assign = fmt.Sprintf("%s: ChatID{ID: %s}", sf.Name, a.Name)
			default:
				a.Assign = fmt.Sprintf("%s: %s", sf.Name, a.Name)
			}
			m.Args = append(m.Args, a)
			continue
		}
		g.addOption(m, f, sf, isText, pm)
	}
	m.Return, m.Result = g.returns(am)
	g.methods = append(g.methods, m)
}

func (g *generator) addOption(m *method, f apiField, sf structField, isText bool, pm string) {
	if f.Name == "reply_markup" && sf.Type.Expr == "*InlineKeyboardMarkup" {
		return // covered by WithInlineKeyboard
	}
	o := &option{
		Name:    sf.Name,
		ArgName: unexported(f.Name),
		ArgType: sf.Type.Expr,
		Doc:     sf.Doc,
		JSON:    f.Name,
	}
	if name, ok := optionNames[m.API.Name+"."+f.Name]; ok {
		o.Name = name
	}
	set := fmt.Sprintf("r.%s = o.v", sf.Name)
	switch {
	case isText:
		o.ArgType = "TextData"
		set = fmt.Sprintf("r.%s, r.%s = o.v.Text, o.v.ParseMode", sf.Name, pm)
	case sf.Type.Expr == "bool":
		o.ArgType = ""
		set = fmt.Sprintf("r.%s = true", sf.Name)
	case sf.Type.Expr == "ChatID":
		o.ArgType = "int64"
		set = fmt.Sprintf("r.%s = ChatID{ID: o.v}", sf.Name)
	case strings.HasPrefix(sf.Type.Expr, "[]"):
		o.Variadic = true
		o.ArgType = strings.TrimPrefix(sf.Type.Expr, "[]")
	}
	g.register(m, o, set)
}

func (g *generator) register(m *method, o *option, set string) {
	have, ok := g.options[o.Name]
	if !ok {
		o.Applies = map[string]string{}
		g.options[o.Name] = o
		have = o
	} else if have.ArgType != o.ArgType || have.Variadic != o.Variadic {
		g.errorf("option With%s takes %q for %s but %q elsewhere; rename one in optionNames",
			o.Name, o.ArgType, m.API.Name, have.ArgType)
		return
	}
	have.Applies[m.Name] = set
	m.Options = append(m.Options, o.Name)
}

func (g *generator) addExtraOptions() {
	for _, e := range extraOptions {
		o := &option{Name: e.Name, Doc: e.Doc, JSON: e.Field}
		if e.Arg != "" {
			o.ArgName, o.ArgType, _ = strings.Cut(e.Arg, " ")
		}
		for _, m := range g.methods {
			if !slices.ContainsFunc(m.Fields, func(f structField) bool { return f.JSON == e.Field }) {
				continue
			}
			g.register(m, o, e.Set)
		}
	}
}

func (g *generator) returns(am *apiMethod) (returnKind, goType) {
	ret := strings.Join(am.Returns, " | ")
	switch ret {
	case "Boolean", "True":
		return returnError, goType{}
	case "Message | Boolean", "Message | True":
		return returnEdit, goType{}
	}
	if len(am.Returns) != 1 {
		g.errorf("%s: unsupported return %q", am.Name, ret)
		return returnError, goType{}
	}
	return returnValue, g.single(am.Returns[0], apiField{Name: "result"}, false)
}

// checkNames fails on generated identifiers that clash.
func (g *generator) checkNames() {
	seen := map[string]string{}
	add := func(name, what string) {
		if prev, ok := seen[name]; ok {
			g.errorf("%s %s clashes with %s", what, name, prev)
		}
		seen[name] = what
	}
	for _, hand := range []string{"Client", "Request", "Error", "TextData", "MessageTarget", "ChatID", "ReplyMarkup", "InputFile", "UpdateType", "MessageType", "UnknownVariant", "ClientOption"} {
		add(hand, "hand-written type")
	}
	for name := range g.structs {
		add(name, "type")
	}
	for name := range g.unions {
		add(name, "union")
	}
	for _, m := range g.methods {
		add(m.Name, "request of "+m.API.Name)
		if len(m.Options) > 0 {
			add(m.Name+"Option", "option interface of "+m.API.Name)
		}
	}
	for name := range g.options {
		add(name+"Option", "option type of "+name)
	}
}
