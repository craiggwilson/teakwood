package slicedict

type options[K comparable, V any] struct {
	initialCapacity int
}

type Opt[K comparable, V any] func(*options[K, V])

func WithInitialCapacity[K comparable, V any](capacity int) Opt[K, V] {
	return func(o *options[K, V]) {
		o.initialCapacity = capacity
	}
}
