package iter

func Limit[T any](it Iter[T], limit int) Iter[T] {
	return &limitIter[T]{
		it:    it,
		limit: limit,
	}
}

type limitIter[T any] struct {
	it    Iter[T]
	limit int

	count int
}

func (it *limitIter[T]) Next() (T, bool) {
	if it.count >= it.limit {
		var def T
		return def, false
	}

	it.count++
	return it.it.Next()
}

func (it *limitIter[T]) Close() error {
	return it.it.Close()
}
