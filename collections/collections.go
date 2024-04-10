package collections

// Sequence represent ordered sequences of elements with common
// operations in a functional way.
type Sequence[T any] interface {
	hasNext() bool
	next() int
	dryCheck()
}

var _ Sequence[any] = &Slice[any]{}

type Predicate[T any] func(T) bool

type Slice[T any] struct {
	dry        bool
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
	return s.index < len(s.idxx)
}

func (s *Slice[T]) next() int {
	s.index++
	return s.idxx[s.index-1]
}

func (s *Slice[T]) dryCheck() {
	s.dry = len(s.idxx) == 0
}

func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {
	if s.dry {
		return s
	}

	var filterIndex int
	for s.hasNext() {
		if p(s.collection[s.next()]) {
			s.idxx[filterIndex] = s.index - 1
			filterIndex++
		}
	}

	s.idxx = s.idxx[:filterIndex]
	s.index = 0

	s.dryCheck()
	return s
}

// Take takes the first n elements of s. If s hast less elements than n, all elements will be taken.
// It panics if n < 0.
func (s *Slice[T]) Take(n int) *Slice[T] {
	return s.takeOrSkip(n, true)
}

// Skip skips the first n elements of s. If s has less elements than n, all elements will be skipped.
// It panics if n < 0.
func (s *Slice[T]) Skip(n int) *Slice[T] {
	return s.takeOrSkip(n, false)
}

// takeOrSkip is a helper function to avoid violating DRY.
// It panics if n < 0.
func (s *Slice[T]) takeOrSkip(n int, take bool) *Slice[T] {
	if s.dry {
		return s
	}

	if len(s.idxx) < n {
		n = len(s.idxx)
	}

	if take {
		s.idxx = s.idxx[:n]
	} else {
		s.idxx = s.idxx[n:]
	}

	s.dryCheck()
	return s
}

func (s *Slice[T]) Collect() []T {
	collection := make([]T, 0, len(s.idxx))
	for _, idx := range s.idxx {
		collection = append(collection, s.collection[idx])
	}
	return collection
}
