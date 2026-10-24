package gogram

// storage is a path-keyed handler lookup backed by a radix tree.
type storage struct {
	tree *tree
}

func (s *storage) Insert(path string, handler Handler) {
	s.tree.Insert(path, handler)
}

func (s *storage) Map() map[string]Handler {
	return s.tree.ToMap()
}

func (s *storage) Find(path string) (Handler, bool) {
	_, handler, ok := s.tree.LongestPrefix(path)
	return handler, ok
}

// newStorage creates a storage.
func newStorage() *storage {
	s := new(storage)
	s.tree = newTree()
	return s
}
