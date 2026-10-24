package gogram

type A struct {
	a int
	b int
}

func (a *A) SetA(v int) {
	a.a = v
}

func (a *A) SetB(v int) {
	a.b = v
}

type B struct {
	a int
}

func (b *B) SetA(v int) {
	b.a = v
}

type ASetter interface {
	SetA(int)
}

type BSetter interface {
	SetB(int)
}

func WithA[T ASetter](v int) func(a T) {
	return func(a T) {
		a.SetA(v)
	}
}

func WithB[T BSetter](v int) func(a T) {
	return func(a T) {
		a.SetB(v)
	}
}

type AOption func(a *A)

func NewA(opts ...AOption) {
	a := &A{}
	for _, o := range opts {
		o(a)
	}
}

func t() {
	NewA(WithA[*A](1))
}
