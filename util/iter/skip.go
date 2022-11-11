package iter

func Skip[T any](it Iter[T], skip int) Iter[T] {
	return &skipIter[T]{
		it:   it,
		skip: skip,
	}
}

type skipIter[T any] struct {
	it   Iter[T]
	skip int

	count int
}

func (it *skipIter[T]) Next() (T, bool) {
	for it.count < it.skip {
		it.count++
		next, ok := it.it.Next()
		if !ok {
			return next, false
		}
	}

	return it.it.Next()
}

func (it *skipIter[T]) Close() error {
	return it.it.Close()
}
