package gogram

// Storage interface
type Storage interface {
	Insert(string, Handler)
	Find(string) (Handler, bool)
	Map() map[string]Handler
}

type storage struct {
	tree *Tree
}

func (s *storage) Insert(path string, handler Handler) {
	s.tree.Insert(path, handler)
}

func (s *storage) Map()  map[string]Handler {
	return s.tree.ToMap()
}

func (s *storage) Find(path string) (Handler, bool) {
	_, handler, ok := s.tree.LongestPrefix(path)
	return handler, ok
}

// NewStorage creates storage
func NewStorage() Storage {
	storage := new(storage)
	storage.tree = NewTree()
	return storage
}