package set

import (
	"github.com/craiggwilson/teakwood/util/collections/set"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ set.ReadOnly[int] = (*FrozenSet[int])(nil)

func New[T comparable](s set.ReadOnly[T]) *FrozenSet[T] {
	return &FrozenSet[T]{s}
}

type FrozenSet[T comparable] struct {
	s set.ReadOnly[T]
}

func (s *FrozenSet[T]) Contains(v T) bool {
	return s.s.Contains(v)
}

func (s *FrozenSet[T]) Iter() iter.Iter[T] {
	return s.s.Iter()
}

func (s *FrozenSet[T]) Len() int {
	return s.s.Len()
}
