package list

import "github.com/craiggwilson/teakwood/util/iter"

func NewSlice[T any](values ...T) *Slice[T] {
	return &Slice[T]{values}
}

func NewSliceWithCap[T any](cap int, values ...T) *Slice[T] {
	vs := make([]T, len(values), cap)
	copy(vs, values)
	return &Slice[T]{vs}
}

type Slice[T any] struct {
	values []T
}

func (l *Slice[T]) Add(values ...T) {
	l.values = append(l.values, values...)
}

func (l *Slice[T]) InsertAt(idx int, values ...T) {
	l.values = append(l.values[:idx], append(values, l.values[idx+1:]...)...)
}

func (l *Slice[T]) Iter() iter.Iter[T] {
	return iter.NewSlice[T](l.values...)
}

func (l *Slice[T]) Len() int {
	return len(l.values)
}

func (l *Slice[T]) RemoveAt(idx int) {
	l.values = append(l.values[:idx], l.values[idx+1:]...)
}

func (l *Slice[T]) Reverse() {
	length := len(l.values) - 1
	for i := 0; i < len(l.values)/2; i++ {
		l.values[i], l.values[length-i] = l.values[length-i], l.values[i]
	}
}

func (l *Slice[T]) Value(idx int) T {
	return l.values[idx]
}
