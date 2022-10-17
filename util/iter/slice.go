package iter

func NewSlice[T any](values ...T) *Slice[T] {
	return &Slice[T]{values: values}
}

type Slice[T any] struct {
	values []T
	pos    int
}

func (it *Slice[T]) Next() (T, bool) {
	if it.pos < len(it.values) {
		it.pos++
		return it.values[it.pos-1], true
	}

	var t T
	return t, false
}

func (it *Slice[T]) Err() error {
	return nil
}
