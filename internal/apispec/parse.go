package apispec

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// coreTypes are the primitive types of the spec.
var coreTypes = []string{"String", "Boolean", "Integer", "Float"}

// approvedMultiReturns are methods results that may be one of several
// types: edits return the new message, or True for inline messages.
var approvedMultiReturns = [][]string{{"Message", "Boolean"}}

// Parse extracts the spec from a snapshot made by Extract.
//
// The page lists sections in order: an <h4> naming a method (lowercase)
// or type (capitalized), then paragraphs describing it, a table of its
// fields, and for unions a list of their members. Parse follows the same
// rules as the scraper of github.com/PaulSonOfLars/telegram-bot-api-spec,
// and validates the result the same way.
func Parse(snapshot []byte) (*Spec, error) {
	root, err := parseHTML(snapshot)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}
	content := root.child("div")
	if content == nil || content.attrs["id"] != "dev_page_content" {
		return nil, errors.New("the snapshot doesn't start with the documentation <div>")
	}

	s := &Spec{Methods: map[string]*Method{}, Types: map[string]*Type{}}
	release := content.child("h4")
	version := content.child("p")
	if release == nil || version == nil {
		return nil, errors.New("the page has no release date or version")
	}
	s.ReleaseDate = release.textOf()
	if a := release.find("a"); a != nil {
		s.Changelog = URL + a.attrs["href"]
	}
	s.Version = version.textOf()

	var method *Method
	var typ *Type
	for _, x := range content.children {
		switch x.tag {
		case "h3", "hr":
			method, typ = nil, nil // a new category
		case "h4":
			method, typ = nil, nil
			anchor := x.find("a")
			if anchor == nil {
				continue
			}
			if name := anchor.attrs["name"]; name != "" && strings.Contains(name, "-") {
				continue // not a method or type, e.g. a release date
			}
			name := x.textOf()
			href := URL + anchor.attrs["href"]
			if name == "" {
				continue
			}
			if unicode.IsUpper([]rune(name)[0]) {
				typ = &Type{Name: name, Href: href}
				s.Types[name] = typ
				s.typeOrder = append(s.typeOrder, name)
			} else {
				method = &Method{Name: name, Href: href}
				s.Methods[name] = method
			}
		case "p":
			lines := cleanDescription(x)
			if method != nil {
				method.Description = append(method.Description, lines...)
			} else if typ != nil {
				typ.Description = append(typ.Description, lines...)
			}
		case "table":
			if err := parseFields(x, method, typ); err != nil {
				return nil, err
			}
		case "ul":
			parseList(x, method, typ)
		}
	}

	for _, m := range s.Methods {
		m.Returns = returnTypes(strings.Join(m.Description, "\n"))
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func parseFields(table *node, method *Method, typ *Type) error {
	if method == nil && typ == nil {
		return nil
	}
	rows := table
	if body := table.find("tbody"); body != nil {
		rows = body
	}
	var fields []Field
	for _, tr := range rows.findAll("tr") {
		cells := tr.findAll("td")
		if len(cells) == 0 {
			continue // a header row
		}
		switch {
		case typ != nil && len(cells) == 3:
			desc := strings.Join(cleanDescription(cells[2]), " ")
			fields = append(fields, Field{
				Name:        cells[0].textOf(),
				Types:       cleanType(cells[1].textOf()),
				Required:    !strings.HasPrefix(desc, "Optional. "),
				Description: desc,
			})
		case method != nil && len(cells) == 4:
			fields = append(fields, Field{
				Name:        cells[0].textOf(),
				Types:       cleanType(cells[1].textOf()),
				Required:    cells[2].textOf() == "Yes",
				Description: strings.Join(cleanDescription(cells[3]), " "),
			})
		default:
			name := ""
			if typ != nil {
				name = typ.Name
			} else {
				name = method.Name
			}
			return fmt.Errorf("%s: a table row has %d cells; has the page layout changed?", name, len(cells))
		}
	}
	if typ != nil {
		typ.Fields = fields
	} else {
		method.Fields = fields
	}
	return nil
}

// parseList handles a list after a method or type: for a type, the
// members of a union. The items are added to the description too.
func parseList(ul *node, method *Method, typ *Type) {
	if typ != nil && typ.Name == "InputFile" {
		return // its list isn't about subtypes
	}
	var items []string
	for _, li := range ul.findAll("li") {
		items = append(items, cleanDescription(li)...)
	}
	var bullets []string
	for _, item := range items {
		bullets = append(bullets, "- "+item)
	}
	switch {
	case typ != nil:
		typ.Subtypes = items
		typ.Description = append(typ.Description, bullets...)
	case method != nil:
		method.Description = append(method.Description, bullets...)
	}
}

var (
	reReturn1 = regexp.MustCompile(`(?i).*(?:on success,)([^.]*)`)
	reReturn2 = regexp.MustCompile(`(?i).*(?:returns)([^.]*)(?:on success)?`)
	reReturn3 = regexp.MustCompile(`(?i).*([^.]*)(?:is returned)`)
	reArrayOf = regexp.MustCompile(`(?i)(?:array of )+(\w*)`)
)

// collapseSpace replaces each run of two or more whitespace characters
// (Unicode whitespace, such as the page's non-breaking spaces) with the
// last character of the run.
func collapseSpace(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i := 0; i < len(runes); {
		if !unicode.IsSpace(runes[i]) {
			b.WriteRune(runes[i])
			i++
			continue
		}
		j := i
		for j < len(runes) && unicode.IsSpace(runes[j]) {
			j++
		}
		b.WriteRune(runes[j-1])
		i = j
	}
	return b.String()
}

// returnTypes reads the result type of a method from its description,
// such as "On success, the sent Message is returned".
func returnTypes(description string) []string {
	for _, re := range []*regexp.Regexp{reReturn1, reReturn2, reReturn3} {
		if m := re.FindStringSubmatch(description); m != nil {
			return extractReturn(strings.TrimSpace(m[1]))
		}
	}
	return nil
}

func extractReturn(s string) []string {
	if m := reArrayOf.FindStringSubmatch(s); m != nil {
		var out []string
		for _, t := range cleanType(m[1]) {
			out = append(out, "Array of "+t)
		}
		return out
	}
	var out []string
	for _, word := range strings.Fields(s) {
		if !unicode.IsUpper([]rune(word)[0]) {
			continue
		}
		word = strings.Map(func(r rune) rune {
			if r < 128 && unicode.IsPunct(r) || strings.ContainsRune("$+<=>^`|~", r) {
				return -1
			}
			return r
		}, word)
		out = append(out, cleanType(word)...)
	}
	return out
}

// cleanDescription returns the text of a description element as lines,
// with links spelled out and typography simplified.
func cleanDescription(n *node) []string {
	var b strings.Builder
	var write func(*node)
	write = func(c *node) {
		switch c.tag {
		case "":
			b.WriteString(c.text)
			return
		case "img":
			b.WriteString(c.attrs["alt"])
			return
		case "br":
			b.WriteString("\n")
			return
		case "a":
			if text := c.textOf(); strings.Contains(text, "»") {
				link := c.attrs["href"]
				switch {
				case strings.HasPrefix(link, "#"):
					link = URL + link
				case strings.HasPrefix(link, "/"):
					link = "https://core.telegram.org" + link
				}
				b.WriteString(strings.ReplaceAll(text, " »", ": "+link))
				return
			}
		}
		for _, cc := range c.children {
			write(cc)
		}
	}
	write(n)

	text := collapseSpace(b.String())
	text = strings.NewReplacer(
		"”", `"`, "“", `"`,
		"…", "...",
		"–", "-", "—", "-",
		"’", "'",
	).Replace(text)

	var lines []string
	for line := range strings.SplitSeq(text, "\n") {
		if line = strings.TrimSpace(line); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// cleanType splits a type cell such as "Integer or String" or
// "Array of InputMediaAudio, InputMediaDocument" into spec type names.
func cleanType(t string) []string {
	prefix := ""
	if rest, ok := strings.CutPrefix(t, "Array of "); ok {
		prefix, t = "Array of ", rest
	}
	var out []string
	for _, or := range strings.Split(t, " or ") {
		for _, and := range strings.Split(strings.TrimSpace(or), " and ") {
			for _, part := range strings.Split(strings.TrimSpace(and), ", ") {
				out = append(out, prefix+properType(strings.TrimSpace(part)))
			}
		}
	}
	return out
}

func properType(t string) string {
	switch t {
	case "Messages":
		return "Message"
	case "Float number":
		return "Float"
	case "Int":
		return "Integer"
	case "True", "Bool":
		return "Boolean"
	}
	return t
}

var reEither = regexp.MustCompile(`(?i)it can be either (.*?), or any of the following types:`)
var reEitherPart = regexp.MustCompile(`(?:a|an)\s+(?:(\w*\s+of\s+\w*)|(\w*))`)

// validate checks that every referenced type exists and every method
// has a result, and links unions to their members.
func (s *Spec) validate() error {
	var errs []error
	known := func(t string) bool {
		for {
			rest, ok := strings.CutPrefix(t, "Array of ")
			if !ok {
				break
			}
			t = rest
		}
		_, isType := s.Types[t]
		return isType || slices.Contains(coreTypes, t)
	}

	for _, name := range s.typeOrder { // page order, for SubtypeOf
		t := s.Types[name]
		if len(t.Fields) == 0 {
			description := strings.Join(t.Description, "")
			if len(t.Subtypes) == 0 && !strings.Contains(strings.ToLower(description), "currently holds no information") && name != "InputFile" {
				errs = append(errs, fmt.Errorf("type %s has no fields or subtypes", name))
				continue
			}
			// "Currently, it can be either a String for plain text, an Array
			// of RichText, or any of the following types:"
			if m := reEither.FindStringSubmatch(description); m != nil {
				var extra []string
				for part := range strings.SplitSeq(m[1], ",") {
					if tm := reEitherPart.FindStringSubmatch(strings.TrimSpace(part)); tm != nil {
						extra = append(extra, tm[1]+tm[2])
					}
				}
				t.Subtypes = append(extra, t.Subtypes...)
			}
			for _, st := range t.Subtypes {
				cleaned := strings.TrimPrefix(st, "Array of ")
				switch {
				case slices.Contains(coreTypes, cleaned) || cleaned == name:
				case s.Types[st] != nil:
					s.Types[st].SubtypeOf = append(s.Types[st].SubtypeOf, name)
				default:
					errs = append(errs, fmt.Errorf("type %s has an unknown subtype %s", name, st))
				}
			}
		}
		for _, f := range t.Fields {
			for _, ft := range f.Types {
				if !known(ft) {
					errs = append(errs, fmt.Errorf("%s.%s has an unknown type %s", name, f.Name, ft))
				}
			}
		}
	}

	for _, name := range sortedKeys(s.Methods) {
		m := s.Methods[name]
		if len(m.Returns) == 0 {
			errs = append(errs, fmt.Errorf("method %s: no result type found in its description", name))
			continue
		}
		if len(m.Returns) > 1 && !slices.ContainsFunc(approvedMultiReturns, func(r []string) bool { return slices.Equal(r, m.Returns) }) {
			errs = append(errs, fmt.Errorf("method %s returns several types %v; approve it in approvedMultiReturns", name, m.Returns))
		}
		for _, f := range m.Fields {
			for _, ft := range f.Types {
				if !known(ft) {
					errs = append(errs, fmt.Errorf("%s.%s has an unknown type %s", name, f.Name, ft))
				}
			}
		}
		for _, r := range m.Returns {
			if !known(r) {
				errs = append(errs, fmt.Errorf("method %s returns an unknown type %s", name, r))
			}
		}
	}
	return errors.Join(errs...)
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
