package collectionscach

type itterator[T any] interface {
	hasNext() bool
	next() T
}

var _ itterator[any] = &Slice[any]{}
var _ itterator[any] = &cache[any]{}

type cache[T any] struct {
	index   int
	records []T
}

func newCache[T any](collectionLen int) cache[T] {
	return cache[T]{
		records: make([]T, 0, collectionLen),
	}
}

func (c *cache[T]) hasNext() bool {
	return c.index <= len(c.records)-1
}

func (c *cache[T]) next() T {
	c.index++
	return c.records[c.index-1]
}

func (c *cache[T]) empty() bool {
	return len(c.records) == 0
}

func (c *cache[T]) add(r T) {
	c.records = append(c.records, r)
}

func (c *cache[T]) invalidate(index int) {
	c.records = append(c.records[:index], c.records[index+1:]...)
	c.index--
}

type Predicate[T any] func(T) bool

type Slice[T any] struct {
	index      int
	collection []T
	cache      cache[int]
}

func NewSlice[T any](slice *[]T) *Slice[T] {
	return &Slice[T]{
		collection: *slice,
		cache:      newCache[int](len(*slice)),
	}
}

func (s *Slice[T]) hasNext() bool {
	return s.index < len(s.collection)-1
}

func (s *Slice[T]) next() T {
	s.index++
	return s.collection[s.index-1]
}

func (s *Slice[T]) Filter(p Predicate[T]) *Slice[T] {

	if s.cache.empty() {
		for s.hasNext() {
			if p(s.next()) {
				s.cache.add(s.index - 1)
			}
		}
	} else {
		for s.cache.hasNext() {
			item := s.cache.next()
			if !p(s.collection[item]) {
				s.cache.invalidate(s.cache.index - 1)
			}
		}
	}

	s.index = 0
	s.cache.index = 0
	return s
}

func (s *Slice[T]) Collect() []T {
	collection := make([]T, 0, len(s.cache.records))
	for _, idx := range s.cache.records {
		collection = append(collection, s.collection[idx])
	}
	return collection
}
