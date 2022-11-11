package sliceset

type options[T comparable] struct {
	initialCapacity int
}

type Opt[T comparable] func(*options[T])

func WithInitialCapacity[T comparable](capacity int) Opt[T] {
	return func(o *options[T]) {
		o.initialCapacity = capacity
	}
}
