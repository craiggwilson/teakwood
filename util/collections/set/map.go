package set

import "github.com/craiggwilson/teakwood/util/iter"

var _ Set[int] = (*Map[int])(nil)

func NewMap[T comparable]() *Map[T] {
	return &Map[T]{make(map[T]struct{})}
}

type Map[T comparable] struct {
	values map[T]struct{}
}

func (s *Map[T]) Add(v T) {
	s.values[v] = struct{}{}
}

func (s *Map[T]) Clear() {
	s.values = make(map[T]struct{})
}

func (s *Map[T]) Contains(v T) bool {
	_, ok := s.values[v]
	return ok
}

func (s *Map[T]) Iter() iter.Iter[T] {
	slice := make([]T, 0, len(s.values))
	for k := range s.values {
		slice = append(slice, k)
	}

	return iter.FromSlice[T](slice...)
}

func (s *Map[T]) Len() int {
	return len(s.values)
}

func (s *Map[T]) Remove(v T) {
	delete(s.values, v)
}
