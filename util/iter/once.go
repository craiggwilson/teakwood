package iter

func Once[T any](value T) Iter[T] {
	return &onceIter[T]{value, false}
}

type onceIter[T any] struct {
	value T
	done  bool
}

func (it *onceIter[T]) Next() (T, bool) {
	if !it.done {
		it.done = true
		return it.value, true
	}

	var def T
	return def, false
}

func (it *onceIter[T]) Close() error {
	return nil
}
