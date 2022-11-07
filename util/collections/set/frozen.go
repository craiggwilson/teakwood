package set

import "github.com/craiggwilson/teakwood/util/iter"

var _ ReadOnly[int] = (*Frozen[int])(nil)

func NewFrozen[T comparable](s ReadOnly[T]) *Frozen[T] {
	return &Frozen[T]{s}
}

type Frozen[T comparable] struct {
	s ReadOnly[T]
}

func (s *Frozen[T]) Contains(v T) bool {
	return s.s.Contains(v)
}

func (s *Frozen[T]) Iter() iter.Iter[T] {
	return s.s.Iter()
}

func (s *Frozen[T]) Len() int {
	return s.s.Len()
}
