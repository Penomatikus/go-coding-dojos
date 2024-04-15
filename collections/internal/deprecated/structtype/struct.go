package structtype

// Sequence represent ordered sequences of elements with common
// operations in a functional way.
type Sequence[T any] interface {
	hasNext() bool
	next() T
}

var _ Sequence[any] = &Slice[any]{}

type Predicate[T any] func(T) bool

type Slice[T any] struct {
	index   int
	content []T
}

// Deprecated: Is was WIP and try and error, do not use
func NewSlice[T any](slice []T) *Slice[T] {
	return &Slice[T]{
		index:   0,
		content: slice,
	}
}

func (s *Slice[T]) hasNext() bool {
	return s.index < len(s.content)
}

func (s *Slice[T]) next() T {
	s.index++
	return s.content[s.index-1]
}

func (s *Slice[T]) Find(p Predicate[T]) *Slice[int] {
	find := make([]int, 0)
	for s.hasNext() {
		if p(s.next()) {
			find = append(find, s.index-1)
		}
	}
	return &Slice[int]{
		content: find,
	}
}

func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {
	find := s.Find(p)
	filter := make([]T, 0, len(find.content))
	for find.hasNext() {
		filter = append(filter, s.content[find.next()])

	}
	s.content = filter
	s.index = 0
	return s
}

func (s *Slice[T]) Collect() []T {
	return s.content
}
