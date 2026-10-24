package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comment(buf *bytes.Buffer, indent string, lines ...string) {
	for _, l := range lines {
		for _, part := range strings.Split(l, "\n") {
			if part == "" {
				fmt.Fprintf(buf, "%s//\n", indent)
				continue
			}
			fmt.Fprintf(buf, "%s// %s\n", indent, part)
		}
	}
}

func (g *generator) writeTypes(buf *bytes.Buffer) {
	buf.WriteString("package telegram\n\nimport (\n\"encoding/json\"\n\"fmt\"\n)\n\n")
	buf.WriteString("var _ = fmt.Errorf // used by the decoders of some types\n\n")
	g.writeUpdateTypes(buf)
	for _, name := range sortedKeys(g.structs) {
		s := g.structs[name]
		doc := append([]string(nil), s.Doc...)
		if s.Href != "" {
			doc = append(doc, "", "See "+s.Href+".")
		}
		comment(buf, "", doc...)
		fmt.Fprintf(buf, "type %s struct {\n", name)
		for _, f := range s.Fields {
			if f.Doc != "" {
				comment(buf, "\t", f.Doc)
			}
			fmt.Fprintf(buf, "\t%s %s %s\n", f.Name, f.Type.Expr, f.tag())
		}
		buf.WriteString("}\n\n")
		if s.TagField != "" {
			fmt.Fprintf(buf, "// MarshalJSON adds the %q discriminator to the encoded object.\n", s.TagField)
			fmt.Fprintf(buf, "func (v %s) MarshalJSON() ([]byte, error) {\n\ttype plain %s\n\treturn marshalTagged(%q, %q, plain(v))\n}\n\n",
				name, name, s.TagField, s.TagValue)
		}
		g.writeUnmarshal(buf, s.Name, s.Fields)
	}
}

// writeUnmarshal decodes the union fields of a struct, if it has any.
func (g *generator) writeUnmarshal(buf *bytes.Buffer, name string, fields []structField) {
	var unionFields []structField
	for _, f := range fields {
		if f.Type.isUnion() {
			if !g.unions[f.Type.Union].Decodable {
				return // input-only type: never decoded
			}
			unionFields = append(unionFields, f)
		}
	}
	if len(unionFields) == 0 {
		return
	}
	fmt.Fprintf(buf, "// UnmarshalJSON decodes the interface-typed fields of %s.\n", name)
	fmt.Fprintf(buf, "func (v *%s) UnmarshalJSON(b []byte) error {\n\ttype plain %s\n\tvar raw struct {\n\t\t*plain\n", name, name)
	for _, f := range unionFields {
		fmt.Fprintf(buf, "\t\t%s %sjson.RawMessage `json:%q`\n", f.Name, strings.Repeat("[]", f.Type.Depth), f.JSON)
	}
	buf.WriteString("\t}\n\traw.plain = (*plain)(v)\n\tif err := json.Unmarshal(b, &raw); err != nil {\n\t\treturn err\n\t}\n")
	for _, f := range unionFields {
		decode := "decode" + f.Type.Union
		switch f.Type.Depth {
		case 0:
			fmt.Fprintf(buf, "\t{\n\t\tx, err := %s(raw.%s)\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"%s: %%w\", err)\n\t\t}\n\t\tv.%s = x\n\t}\n",
				decode, f.Name, f.JSON, f.Name)
		case 1:
			fmt.Fprintf(buf, "\tv.%s = nil\n\tfor _, r := range raw.%s {\n\t\tx, err := %s(r)\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"%s: %%w\", err)\n\t\t}\n\t\tv.%s = append(v.%s, x)\n\t}\n",
				f.Name, f.Name, decode, f.JSON, f.Name, f.Name)
		case 2:
			fmt.Fprintf(buf, "\tv.%s = nil\n\tfor _, row := range raw.%s {\n\t\tvar xs []%s\n\t\tfor _, r := range row {\n\t\t\tx, err := %s(r)\n\t\t\tif err != nil {\n\t\t\t\treturn fmt.Errorf(\"%s: %%w\", err)\n\t\t\t}\n\t\t\txs = append(xs, x)\n\t\t}\n\t\tv.%s = append(v.%s, xs)\n\t}\n",
				f.Name, f.Name, f.Type.Union, decode, f.JSON, f.Name, f.Name)
		default:
			g.errorf("%s.%s: unions nested %d deep are not supported", name, f.JSON, f.Type.Depth)
		}
	}
	buf.WriteString("\treturn nil\n}\n\n")
}

func (g *generator) writeUpdateTypes(buf *bytes.Buffer) {
	u := g.structs["Update"]
	buf.WriteString("// Update types, as used by Update.Type and the allowed_updates parameters.\nconst (\n")
	for _, f := range u.Fields[1:] {
		fmt.Fprintf(buf, "\tUpdate%s UpdateType = %q\n", f.Name, f.JSON)
	}
	buf.WriteString(")\n\n")
	buf.WriteString("// AllUpdateTypes lists every update type, for WithAllowedUpdates.\n")
	buf.WriteString("var AllUpdateTypes = []UpdateType{\n")
	for _, f := range u.Fields[1:] {
		fmt.Fprintf(buf, "\tUpdate%s,\n", f.Name)
	}
	buf.WriteString("}\n\n")
	buf.WriteString("// Type returns the kind of the update, or \"\" for kinds this version\n// of the package doesn't know.\nfunc (u *Update) Type() UpdateType {\n\tswitch {\n")
	for _, f := range u.Fields[1:] {
		fmt.Fprintf(buf, "\tcase u.%s != nil:\n\t\treturn Update%s\n", f.Name, f.Name)
	}
	buf.WriteString("\tdefault:\n\t\treturn \"\"\n\t}\n}\n\n")

	var names []string
	for _, f := range u.Fields[1:] {
		if f.Type.Expr == "*Message" {
			names = append(names, f.Name)
		}
	}
	fmt.Fprintf(buf, "// AnyMessage returns the message the update carries: whichever of\n// %s is set, or nil.\n", strings.Join(names, ", "))
	buf.WriteString("func (u *Update) AnyMessage() *Message {\n\tswitch {\n")
	for _, n := range names {
		fmt.Fprintf(buf, "\tcase u.%s != nil:\n\t\treturn u.%s\n", n, n)
	}
	buf.WriteString("\tdefault:\n\t\treturn nil\n\t}\n}\n\n")
}

func (g *generator) writeUnions(buf *bytes.Buffer) {
	buf.WriteString("package telegram\n\nimport (\n\"bytes\"\n\"encoding/json\"\n)\n\n")
	buf.WriteString("var _ = bytes.TrimSpace // used by the RichText decoder\n\n")
	for _, name := range sortedKeys(g.unions) {
		u := g.unions[name]
		doc := typeDoc(name, u.Doc)
		// The spec lists the members as "- Member" lines; they are repeated below.
		var kept []string
		for _, l := range doc {
			if !strings.HasPrefix(l, "- ") {
				kept = append(kept, l)
			}
		}
		kept = append(kept, "", "Its values are of these types:", "")
		for _, m := range u.Members {
			kept = append(kept, "  - *"+m)
		}
		if u.RichText {
			kept = append(kept, "  - RichTextString", "  - RichTextArray")
		}
		if u.Decodable {
			kept = append(kept, "  - *UnknownVariant, for kinds this version of the package doesn't know")
		}
		if u.Href != "" {
			kept = append(kept, "", "See "+u.Href+".")
		}
		comment(buf, "", kept...)
		fmt.Fprintf(buf, "type %s interface {\n\tis%s()\n}\n\n", name, name)
		for _, m := range u.Members {
			fmt.Fprintf(buf, "func (*%s) is%s() {}\n", m, name)
		}
		if u.RichText {
			fmt.Fprintf(buf, "func (RichTextString) is%s() {}\n", name)
			fmt.Fprintf(buf, "func (RichTextArray) is%s() {}\n", name)
		}
		if u.Decodable {
			fmt.Fprintf(buf, "func (*UnknownVariant) is%s() {}\n", name)
		}
		buf.WriteString("\n")
		if u.Decodable {
			g.writeDecoder(buf, u)
		}
	}
}

func (g *generator) writeDecoder(buf *bytes.Buffer, u *union) {
	fmt.Fprintf(buf, "func decode%s(b []byte) (%s, error) {\n", u.Name, u.Name)
	buf.WriteString("\tif len(b) == 0 || string(b) == \"null\" {\n\t\treturn nil, nil\n\t}\n")
	if u.RichText {
		buf.WriteString("\tswitch bytes.TrimSpace(b)[0] {\n")
		buf.WriteString("\tcase '\"':\n\t\tvar s RichTextString\n\t\terr := json.Unmarshal(b, &s)\n\t\treturn s, err\n")
		fmt.Fprintf(buf, "\tcase '[':\n\t\tvar raws []json.RawMessage\n\t\tif err := json.Unmarshal(b, &raws); err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tarr := make(RichTextArray, 0, len(raws))\n\t\tfor _, r := range raws {\n\t\t\tx, err := decode%s(r)\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\tarr = append(arr, x)\n\t\t}\n\t\treturn arr, nil\n\t}\n", u.Name)
	}
	fmt.Fprintf(buf, "\tvar tag struct {\n\t\tValue string `json:%q`\n\t}\n", u.TagField)
	buf.WriteString("\tif err := json.Unmarshal(b, &tag); err != nil {\n\t\treturn nil, err\n\t}\n")
	fmt.Fprintf(buf, "\tvar v %s\n\tswitch tag.Value {\n", u.Name)
	for _, m := range u.Members {
		fmt.Fprintf(buf, "\tcase %q:\n\t\tv = new(%s)\n", u.Tags[m], m)
	}
	buf.WriteString("\tdefault:\n\t\treturn &UnknownVariant{Raw: append(json.RawMessage(nil), b...)}, nil\n\t}\n")
	buf.WriteString("\tif err := json.Unmarshal(b, v); err != nil {\n\t\treturn nil, err\n\t}\n\treturn v, nil\n}\n\n")
}
