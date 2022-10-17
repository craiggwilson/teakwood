package set

import "github.com/craiggwilson/teakwood/util/iter"

func NewFrozen[T comparable](s ReadOnly[T]) *Frozen[T] {
	return &Frozen[T]{s}
}

type Frozen[T comparable] struct {
	s ReadOnly[T]
}

func (s *Frozen[T]) Contains(value T) bool {
	return s.s.Contains(value)
}

func (s *Frozen[T]) Iter() iter.Iter[T] {
	return s.s.Iter()
}

func (s *Frozen[T]) Len() int {
	return s.s.Len()
}
