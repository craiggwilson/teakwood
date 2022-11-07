package set

import (
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ Set[int] = (*Slice[int])(nil)

func NewSlice[T comparable]() *Slice[T] {
	return &Slice[T]{}
}

type Slice[T comparable] struct {
	values []T
}

func (s *Slice[T]) Add(v T) {
	if s.Contains(v) {
		return
	}

	s.values = append(s.values, v)
}

func (s *Slice[T]) Clear() {
	s.values = s.values[:0]
}

func (s *Slice[T]) Contains(v T) bool {
	for _, sv := range s.values {
		if v == sv {
			return true
		}
	}

	return false
}

func (s *Slice[T]) Iter() iter.Iter[T] {
	return iter.FromSlice[T](s.values)
}

func (s *Slice[T]) Len() int {
	return len(s.values)
}

func (s *Slice[T]) Remove(v T) {
	for i, sv := range s.values {
		if v == sv {
			s.values[i] = s.values[len(s.values)-1]
			s.values = s.values[:len(s.values)]
			return
		}
	}
}
