package slice

import "github.com/Penomatikus/collections/internal"

var _ internal.Sequence[any] = &Slice[any]{}

type Predicate[T any] func(T) bool
type Reducer[T any] func(T, T) T

type Slice[T any] struct {
	dry        bool
	index      int
	collection []T
	idxx       []int
}

func New[T any](slice []T) *Slice[T] {
	s := &Slice[T]{
		collection: slice,
		idxx:       make([]int, 0, len(slice)),
	}
	for i := 0; i < len(slice); i++ {
		s.idxx = append(s.idxx, i)
	}

	return s
}

func (s *Slice[T]) HasNext() bool {
	return s.index < len(s.idxx)
}

func (s *Slice[T]) Next() int {
	s.index++
	return s.idxx[s.index-1]
}

func (s *Slice[T]) DryCheck() {
	s.dry = len(s.idxx) == 0
}

// Filter applies p to s:
//	- Depending on p, Filter migth dry s.
//	- If s is dry s is returned.
func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {
	if s.dry {
		return s
	}

	var filterIndex int
	for s.HasNext() {
		next := s.Next()
		if p(s.collection[next]) {
			s.idxx[filterIndex] = next
			filterIndex++
		}
	}

	s.idxx = s.idxx[:filterIndex]
	s.index = 0

	s.DryCheck()
	return s
}

// Take takes the first n elements of s:
//	- If s hast less elements than n, all elements will be taken.
// 	- Depending on n, Take migth dry s.
//	- If s is dry, s is returned.
//	- It panics if n < 0.
func (s *Slice[T]) Take(n int) *Slice[T] {
	return s.takeOrSkip(n, true)
}

// Skip skips the first n elements of s:
//	- If s has less elements than n, all elements will be skipped.
// 	- Depending on n, Skip migth dry s.
//	- If s is dry, s is returned.
//	- It panics if n < 0.
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

	s.DryCheck()
	return s
}

// Collect collects the allocated collection of s.
// 	- It will not dry s
// 	- If s is dry tt is go default
func (s *Slice[T]) Collect() (tt []T) {
	tt = make([]T, 0, len(s.idxx))

	if s.dry {
		return tt
	}

	for s.HasNext() {
		tt = append(tt, s.collection[s.Next()])
	}

	return tt
}

// Reduce applies r to to s.
//	- It will not dry s or change its current index.
// 	- If s is dry t is go default
func (s *Slice[T]) Reduce(r Reducer[T]) (t T) {
	if s.dry {
		return t
	}

	for s.HasNext() {
		t = r(t, s.collection[s.Next()])
	}

	return t
}
