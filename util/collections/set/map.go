package set

import "github.com/craiggwilson/teakwood/util/iter"

func NewMap[T comparable](values ...T) *Map[T] {
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return &Map[T]{m}
}

type Map[T comparable] struct {
	values map[T]struct{}
}

func (s *Map[T]) Add(value T) {
	s.values[value] = struct{}{}
}

func (s *Map[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

func (s *Map[T]) Iter() iter.Iter[T] {
	return iter.NewSlice[T](s.Slice()...)
}

func (s *Map[T]) Len() int {
	return len(s.values)
}

func (s *Map[T]) Remove(value T) {
	delete(s.values, value)
}

func (s *Map[T]) Slice() []T {
	slice := make([]T, 0, len(s.values))
	for k := range s.values {
		slice = append(slice, k)
	}

	return slice
}
