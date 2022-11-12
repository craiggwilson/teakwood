package iter

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

func FromMap[K comparable, V any](m map[K]V) Iterer[KeyValuePair[K, V]] {
	values := make([]KeyValuePair[K, V], 0, len(m))
	for k, v := range m {
		values = append(values, KeyValuePair[K, V]{k, v})
	}

	return FromSlice(values)
}

func FromSlice[T any](values []T) Iterer[T] {
	return &sliceIterer[T]{
		values: values,
	}
}

type sliceIterer[T any] struct {
	values []T
}

func (itr *sliceIterer[T]) Iter() Iter[T] {
	return &sliceIter[T]{
		values: itr.values,
		pos:    0,
	}
}

func (itr *sliceIterer[T]) ToSlice() []T {
	return itr.values
}

type sliceIter[T any] struct {
	values []T
	pos    int
}

func (it *sliceIter[T]) Next() (T, bool) {
	if it.pos < len(it.values) {
		it.pos++
		return it.values[it.pos-1], true
	}

	var t T
	return t, false
}

func (it *sliceIter[T]) Close() error {
	return nil
}

func (it *sliceIter[T]) ToSlice() []T {
	cpy := make([]T, len(it.values)-it.pos)
	copy(cpy, it.values[it.pos:])
	it.pos = len(it.values)
	return cpy
}
