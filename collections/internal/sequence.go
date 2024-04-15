package internal

// Sequence represent ordered sequences of elements with common
// operations in a functional way.
type Sequence[T any] interface {
	HasNext() bool
	Next() int
	DryCheck()
}
