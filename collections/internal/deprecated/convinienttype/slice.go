package convinienttype

type Slice[T any] []T

type Predicate[T any] func(T) bool
type Mapper[T any] func(T) T

// Deprecated: Is was WIP and try and error, do not use
func New[T any](elems []T) *Slice[T] {
	s := append(make(Slice[T], 0, len(elems)), elems...)
	return &s
}

func (s *Slice[T]) Map(m Mapper[T]) *Slice[T] {
	mapped := make(Slice[T], 0, len(*s))
	for _, elem := range *s {
		mapped = append(mapped, m(elem))
	}
	return &mapped
}

func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {
	filter := (*s)[:0]
	for _, elem := range *s {
		if p(elem) {
			filter = append(filter, elem)
		}
	}
	return &filter
}

func (s *Slice[T]) First(p Predicate[T]) *T {
	for _, elem := range *s {
		if p(elem) {
			return &elem
		}
	}
	return nil
}

func (s *Slice[T]) Take(take int) *Slice[T] {
	if take < 0 {
		panic("Cant take negative amount")
	}
	takes := make(Slice[T], 0, len(*s))
	for i := 0; i <= take; i++ {
		takes = append(takes, (*s)[i])
	}
	return &takes
}

func (s *Slice[T]) Skip(skip int) *Slice[T] {
	if skip < 0 {
		panic("Cant skip negative amount")
	}
	after := make(Slice[T], 0, len(*s))
	for i := skip; i <= len(*s); i++ {
		after = append(after, (*s)[i])
	}
	return &after
}

// Collect returns a working copy of the underling array of s
func (s *Slice[T]) Collect() []T {
	release := make([]T, 0, len(*s))
	return append(release, (*s)...)
}
