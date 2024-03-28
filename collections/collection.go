package collections

// Sequence represent ordered sequences of elements with common
// operations in a functional way.
type Sequence[T any] interface {
	hasNext() bool
	next() int
}

var _ Sequence[any] = &Slice[any]{}

type Predicate[T any] func(T) bool

type Slice[T any] struct {
	index      int
	collection []T
	idxx       []int
}

func NewSlice[T any](slice []T) *Slice[T] {
	s := &Slice[T]{
		collection: slice,
		idxx:       make([]int, 0, len(slice)),
	}
	for i := 0; i < len(slice); i++ {
		s.idxx = append(s.idxx, i)
	}

	return s
}

func (s *Slice[T]) hasNext() bool {
	return s.index < len(s.idxx)-1
}

func (s *Slice[T]) next() int {
	s.index++
	return s.idxx[s.index]
}

func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {
	filter := make([]int, 0, len(s.idxx))
	for s.hasNext() {
		if p(s.collection[s.next()]) {
			filter = append(filter, s.index)
		}
	}
	s.idxx = filter
	s.index = 0
	return s
}

func (s *Slice[T]) Collect() []T {
	collection := make([]T, 0, len(s.idxx))
	for _, idx := range s.idxx {
		collection = append(collection, s.collection[idx])
	}
	return collection
}
