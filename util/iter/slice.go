package iter

func FromSlice[T any](values []T) Iter[T] {
	return &sliceIter[T]{values: values}
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
	return cpy
}

func ToSlice[T any](it Iter[T]) []T {
	if slicer, ok := it.(interface{ ToSlice() []T }); ok {
		return slicer.ToSlice()
	}

	var results []T
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		results = append(results, v)
	}

	return results
}
