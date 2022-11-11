package iter

func Repeat[T any](e T, count int) Iter[T] {
	return &repeatIter[T]{
		e:     e,
		count: count,
	}
}

type repeatIter[T any] struct {
	e     T
	count int

	cur int
}

func (it *repeatIter[T]) Next() (T, bool) {
	if it.cur >= it.count {
		var def T
		return def, false
	}

	it.cur++
	return it.e, true
}

func (it *repeatIter[T]) Close() error {
	return nil
}
