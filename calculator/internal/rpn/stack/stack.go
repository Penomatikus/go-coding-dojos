package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("stack is already empty")
	}

	i := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return i
}

// Peek
func (s *Stack[T]) Peek() (T, bool) {
	if s.Empty() {
		var e T
		return e, false
	}

	return s.items[s.Len()-1], true
}

func (s *Stack[T]) PeekAll() []T {
	return s.items
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
