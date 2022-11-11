package iter

import "golang.org/x/exp/constraints"

func Range[T constraints.Integer | constraints.Float](from T, to T, step T) Iter[T] {
	return &rangeIter[T]{
		from: from,
		to:   to,
		step: step,
	}
}

func RangeInclusive[T constraints.Integer | constraints.Float](from T, to T, step T) Iter[T] {
	return &rangeIter[T]{
		from:      from,
		to:        to,
		step:      step,
		inclusive: true,
	}
}

type rangeIter[T constraints.Integer | constraints.Float] struct {
	count     int
	from      T
	to        T
	step      T
	inclusive bool
}

func (it *rangeIter[T]) Next() (T, bool) {
	next := T(it.count)*it.step + it.from
	it.count++
	if next < it.to || (next <= it.to && it.inclusive) {
		return next, true
	}

	var def T
	return def, false
}

func (it *rangeIter[T]) Close() error {
	return nil
}
