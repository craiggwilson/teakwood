package slicelist

type options[T any] struct {
	initialCapacity int
}

type Opt[T any] func(*options[T])

func WithInitialCapacity[T any](capacity int) Opt[T] {
	return func(o *options[T]) {
		o.initialCapacity = capacity
	}
}
