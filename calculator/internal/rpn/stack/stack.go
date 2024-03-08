package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Pop() (*T, bool) {
	if s.Empty() {
		return nil, false
	}

	i := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return &i, true
}

func (s *Stack[T]) Peek() (*T, bool) {
	if s.Empty() {
		return nil, false
	}

	i := s.items[s.Len()-1]
	return &i, true
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
