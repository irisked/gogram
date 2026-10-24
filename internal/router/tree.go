// Package router implements a path-keyed lookup structure (a radix
// tree) used to dispatch an incoming update to the handler registered
// for its path. It has no knowledge of Telegram types -- the value
// type is generic so this package never needs to import, and can
// never be imported by, the root gogram package.
package router

import (
	"sort"
	"strings"
)

// leafNode holds the value stored at a path.
type leafNode[V any] struct {
	key string
	val V
}

// edge is an edge to a child node, keyed by its first byte.
type edge[V any] struct {
	label byte
	node  *node[V]
}

type node[V any] struct {
	leaf *leafNode[V]

	// prefix is the common prefix this node represents.
	prefix string

	// edges are kept sorted by label for binary search.
	edges edges[V]
}

func (n *node[V]) isLeaf() bool {
	return n.leaf != nil
}

func (n *node[V]) addEdge(e edge[V]) {
	n.edges = append(n.edges, e)
	n.edges.Sort()
}

func (n *node[V]) updateEdge(label byte, child *node[V]) {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		n.edges[idx].node = child
		return
	}
	panic("router: replacing missing edge")
}

func (n *node[V]) getEdge(label byte) *node[V] {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		return n.edges[idx].node
	}
	return nil
}

type edges[V any] []edge[V]

func (e edges[V]) Len() int           { return len(e) }
func (e edges[V]) Less(i, j int) bool { return e[i].label < e[j].label }
func (e edges[V]) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e edges[V]) Sort()              { sort.Sort(e) }

// Tree is a radix tree keyed by string paths, supporting exact and
// longest-prefix lookups.
type Tree[V any] struct {
	root *node[V]
}

// New returns an empty Tree.
func New[V any]() *Tree[V] {
	return &Tree[V]{root: &node[V]{}}
}

// longestCommonPrefix returns the length of the shared prefix of two
// strings.
func longestCommonPrefix(k1, k2 string) int {
	max := len(k1)
	if l := len(k2); l < max {
		max = l
	}
	var i int
	for i = 0; i < max; i++ {
		if k1[i] != k2[i] {
			break
		}
	}
	return i
}

// Insert adds a new entry or updates an existing one. Returns the
// previous value and whether it was an update.
func (t *Tree[V]) Insert(s string, v V) (V, bool) {
	var zero V
	var parent *node[V]
	n := t.root
	search := s
	for {
		if len(search) == 0 {
			if n.isLeaf() {
				old := n.leaf.val
				n.leaf.val = v
				return old, true
			}
			n.leaf = &leafNode[V]{key: s, val: v}
			return zero, false
		}

		parent = n
		n = n.getEdge(search[0])

		if n == nil {
			e := edge[V]{
				label: search[0],
				node: &node[V]{
					leaf:   &leafNode[V]{key: s, val: v},
					prefix: search,
				},
			}
			parent.addEdge(e)
			return zero, false
		}

		commonPrefix := longestCommonPrefix(search, n.prefix)
		if commonPrefix == len(n.prefix) {
			search = search[commonPrefix:]
			continue
		}

		child := &node[V]{prefix: search[:commonPrefix]}
		parent.updateEdge(search[0], child)

		child.addEdge(edge[V]{label: n.prefix[commonPrefix], node: n})
		n.prefix = n.prefix[commonPrefix:]

		leaf := &leafNode[V]{key: s, val: v}
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.leaf = leaf
			return zero, false
		}

		child.addEdge(edge[V]{
			label: search[0],
			node:  &node[V]{leaf: leaf, prefix: search},
		})
		return zero, false
	}
}

// Get looks up a key, returning the value and whether it was found.
func (t *Tree[V]) Get(s string) (V, bool) {
	var zero V
	n := t.root
	search := s
	for {
		if len(search) == 0 {
			if n.isLeaf() {
				return n.leaf.val, true
			}
			break
		}
		n = n.getEdge(search[0])
		if n == nil {
			break
		}
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	return zero, false
}

// LongestPrefix is like Get, but returns the longest prefix match
// instead of requiring an exact one.
func (t *Tree[V]) LongestPrefix(s string) (string, V, bool) {
	var zero V
	var last *leafNode[V]
	n := t.root
	search := s
	for {
		if n.isLeaf() {
			last = n.leaf
		}
		if len(search) == 0 {
			break
		}
		n = n.getEdge(search[0])
		if n == nil {
			break
		}
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	if last != nil {
		return last.key, last.val, true
	}
	return "", zero, false
}
