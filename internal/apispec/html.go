package apispec

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

const contentStart = `<div id="dev_page_content">`

// Extract returns the documentation part of the Bot API page: its
// <div id="dev_page_content">. The rest of the page (navigation, scripts,
// and a "page generated in" comment that changes on every request) is
// dropped, so the snapshot changes only when the documentation does.
func Extract(page []byte) ([]byte, error) {
	start := bytes.Index(page, []byte(contentStart))
	if start < 0 {
		return nil, errors.New("the page has no " + contentStart + "; has its layout changed?")
	}
	depth, i := 0, start
	for {
		open := bytes.Index(page[i:], []byte("<div"))
		closing := bytes.Index(page[i:], []byte("</div>"))
		switch {
		case closing < 0:
			return nil, errors.New("the documentation <div> is never closed")
		case open >= 0 && open < closing:
			depth++
			i += open + len("<div")
		default:
			depth--
			i += closing + len("</div>")
			if depth == 0 {
				out := append([]byte(nil), page[start:i]...)
				return append(out, '\n'), nil
			}
		}
	}
}

// node is an element or, with an empty tag, a text node.
type node struct {
	tag      string
	attrs    map[string]string
	children []*node
	text     string
}

// parseHTML parses an HTML fragment. The Bot API page is regular enough
// for encoding/xml in its lenient HTML mode.
func parseHTML(fragment []byte) (*node, error) {
	d := xml.NewDecoder(bytes.NewReader(fragment))
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity

	root := &node{tag: "#root"}
	stack := []*node{root}
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		top := stack[len(stack)-1]
		switch t := tok.(type) {
		case xml.StartElement:
			n := &node{tag: strings.ToLower(t.Name.Local), attrs: map[string]string{}}
			for _, a := range t.Attr {
				n.attrs[strings.ToLower(a.Name.Local)] = a.Value
			}
			top.children = append(top.children, n)
			stack = append(stack, n)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			top.children = append(top.children, &node{text: string(t)})
		}
	}
	return root, nil
}

// textOf returns the text of n and its descendants.
func (n *node) textOf() string {
	var b strings.Builder
	n.walk(func(c *node) bool {
		if c.tag == "" {
			b.WriteString(c.text)
		}
		return true
	})
	return b.String()
}

// walk calls fn for n and its descendants, depth first, skipping the
// descendants of nodes for which fn returns false.
func (n *node) walk(fn func(*node) bool) {
	if !fn(n) {
		return
	}
	for _, c := range n.children {
		c.walk(fn)
	}
}

// find returns the first descendant with the given tag, or nil.
func (n *node) find(tag string) *node {
	var found *node
	n.walk(func(c *node) bool {
		if found != nil {
			return false
		}
		if c != n && c.tag == tag {
			found = c
			return false
		}
		return true
	})
	return found
}

// findAll returns the descendants with the given tag.
func (n *node) findAll(tag string) []*node {
	var out []*node
	n.walk(func(c *node) bool {
		if c != n && c.tag == tag {
			out = append(out, c)
		}
		return true
	})
	return out
}

// child returns the first child element with the given tag, or nil.
func (n *node) child(tag string) *node {
	for _, c := range n.children {
		if c.tag == tag {
			return c
		}
	}
	return nil
}
