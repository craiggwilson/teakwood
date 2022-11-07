package iter

func Filter[T any](it Iter[T], filter func(T) bool) Iter[T] {
	return &filterIter[T]{it, filter}
}

type filterIter[T any] struct {
	it     Iter[T]
	filter func(T) bool
}

func (it *filterIter[T]) Next() (T, bool) {
	for {
		n, ok := it.it.Next()
		if !ok {
			return n, false
		}

		if it.filter(n) {
			return n, true
		}
	}
}

func (it *filterIter[T]) Close() error {
	return it.it.Close()
}
