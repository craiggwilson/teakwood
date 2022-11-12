package iter

func Append[S any](dst Iterer[S], values ...S) Iterer[S] {
	return Concat(dst, FromSlice(values))
}

func Chunk[S any](src Iterer[S], size int) Iterer[[]S] {
	panic("not implemented")
}

func Concat[S any](first Iterer[S], second Iterer[S]) Iterer[S] {
	return ItererFunc[S](func() Iter[S] {
		return &concatIter[S]{
			first:  first.Iter(),
			second: second.Iter(),
		}
	})
}

type concatIter[S any] struct {
	first  Iter[S]
	second Iter[S]

	firstDone bool
}

func (it *concatIter[S]) Next() (S, bool) {
	if !it.firstDone {
		value, ok := it.first.Next()
		if ok {
			return value, ok
		}

		it.firstDone = true
	}

	return it.second.Next()
}

func (it *concatIter[S]) Close() error {
	err1 := it.first.Close()
	err2 := it.second.Close()

	if err1 != nil {
		return err1
	}

	return err2
}

func Distinct[S comparable](src Iterer[S]) Iterer[S] {
	panic("not implemented")
}

func DistinctBy[S any, K comparable](src Iterer[S], mapper func(S) K) Iterer[S] {
	panic("not implemented")
}

func Filter[S any](src Iterer[S], filter func(S) bool) Iterer[S] {
	return ItererFunc[S](func() Iter[S] {
		return &filterIter[S]{
			src:    src.Iter(),
			filter: filter,
		}
	})
}

type filterIter[S any] struct {
	src    Iter[S]
	filter func(S) bool
}

func (it *filterIter[S]) Next() (S, bool) {
	for {
		elem, ok := it.src.Next()
		if !ok {
			return elem, false
		}

		if it.filter(elem) {
			return elem, true
		}
	}
}

func (it *filterIter[S]) Close() error {
	return it.src.Close()
}

type Grouping[S any, K comparable] struct {
	Key  K
	Iter Iter[S]
}

func Group[S any, K comparable](src Iterer[S], mapper func(S) K) Iterer[Grouping[S, K]] {
	panic("not implemented")
}

func Map[S, R any](src Iterer[S], mapper func(S) R) Iterer[R] {
	return ItererFunc[R](func() Iter[R] {
		return &mapIter[S, R]{
			src:    src.Iter(),
			mapper: mapper,
		}
	})
}

type mapIter[S, R any] struct {
	src    Iter[S]
	mapper func(S) R
}

func (it *mapIter[S, R]) Next() (R, bool) {
	value, ok := it.src.Next()
	if !ok {
		var def R
		return def, false
	}

	return it.mapper(value), true
}

func (it *mapIter[S, R]) Close() error {
	return it.src.Close()
}

func Prepend[S any](src Iterer[S], values ...S) Iterer[S] {
	return Concat(FromSlice(values), src)
}

func SelectMany[S, R any](src Iterer[S], selector func(S) Iterer[R]) Iterer[R] {
	panic("not implemented")
}

func Skip[S any](src Iterer[S], skip int) Iterer[S] {
	return ItererFunc[S](func() Iter[S] {
		return &skipIter[S]{
			src:  src.Iter(),
			skip: skip,
		}
	})
}

type skipIter[S any] struct {
	src  Iter[S]
	skip int

	count int
}

func (it *skipIter[S]) Next() (S, bool) {
	for it.count < it.skip {
		it.count++
		value, ok := it.src.Next()
		if !ok {
			return value, false
		}
	}

	return it.src.Next()
}

func (it *skipIter[S]) Close() error {
	return it.src.Close()
}

func Take[S any](src Iterer[S], limit int) Iterer[S] {
	return ItererFunc[S](func() Iter[S] {
		return &takeIter[S]{
			src:   src.Iter(),
			limit: limit,
		}
	})
}

type takeIter[S any] struct {
	src   Iter[S]
	limit int

	count int
}

func (it *takeIter[S]) Next() (S, bool) {
	if it.count >= it.limit {
		var def S
		return def, false
	}

	it.count++
	return it.src.Next()
}

func (it *takeIter[S]) Close() error {
	return it.src.Close()
}

func Zip[S1, S2, R any](first Iterer[S1], second Iterer[S2], zipper func(S1, S2) R) Iterer[R] {
	return ItererFunc[R](func() Iter[R] {
		return &zipIter[S1, S2, R]{
			first:  first.Iter(),
			second: second.Iter(),
			zipper: zipper,
		}
	})
}

type zipIter[S1, S2, R any] struct {
	first  Iter[S1]
	second Iter[S2]

	zipper func(S1, S2) R
}

func (it *zipIter[S1, S2, R]) Next() (R, bool) {
	value1, ok1 := it.first.Next()
	value2, ok2 := it.second.Next()

	if ok1 && ok2 {
		return it.zipper(value1, value2), true
	}

	var def R
	return def, false
}

func (it *zipIter[S1, S2, R]) Close() error {
	err1 := it.first.Close()
	err2 := it.second.Close()

	if err1 != nil {
		return err1
	}

	return err2
}
