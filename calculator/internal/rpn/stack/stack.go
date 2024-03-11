// package stack provides an implementation of the stack datastructure
package stack

type Stack[T any] struct {
	items []T
}

// New creates an emtpy Stack of T
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Len returns the number of elements in s.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// Empty tells if s is empty.
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Pop will release the top element of s or panics, if the stack is empty.
func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("stack is already empty")
	}

	i := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return i
}

// Peek will return the top element of s without releasing it. Returns the default of T and false, if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if s.Empty() {
		var e T
		return e, false
	}

	return s.items[s.Len()-1], true
}

// PeekAll returns the content of s as slice, without releasing any element.
func (s *Stack[T]) PeekAll() []T {
	return s.items
}

// Push adds item to s.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
