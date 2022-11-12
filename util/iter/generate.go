package iter

import "golang.org/x/exp/constraints"

func Generate[T any](generator func() (T, bool)) Iterer[T] {
	return ItererFunc[T](func() Iter[T] {
		return &generateIter[T]{
			generator: generator,
		}
	})
}

type generateIter[T any] struct {
	generator func() (T, bool)
	done      bool
}

func (it *generateIter[T]) Next() (T, bool) {
	if it.done {
		var def T
		return def, false
	}

	value, ok := it.generator()
	if !ok {
		it.done = true
	}

	return value, ok
}

func (it *generateIter[T]) Close() error {
	return nil
}

func Range[T constraints.Integer | constraints.Float](from T, to T, step T) Iterer[T] {
	count := T(0)
	return Generate(func() (T, bool) {
		value := from + count*step
		count++
		if value < to {
			return value, true
		}

		var def T
		return def, false
	})
}

func RangeInclusive[T constraints.Integer | constraints.Float](from T, to T, step T) Iterer[T] {
	count := T(0)
	return Generate(func() (T, bool) {
		value := from + count*step
		count++
		if value <= to {
			return value, true
		}

		var def T
		return def, false
	})
}

func Repeat[T any](value T, count int) Iterer[T] {
	i := 0
	return Generate(func() (T, bool) {
		if i >= count {
			var def T
			return def, false
		}

		i++
		return value, true
	})
}
