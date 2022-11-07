package iter

func Empty[T any]() Iter[T] {
	return emptyIter[T]{}
}

type emptyIter[T any] struct{}

func (emptyIter[T]) Next() (T, bool) {
	var def T
	return def, false
}

func (emptyIter[T]) Close() error {
	return nil
}
