package sliceset

import (
	"github.com/craiggwilson/teakwood/util/collections/set"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ set.Set[int] = (*SliceSet[int])(nil)

func New[T comparable](opts ...Opt[T]) *SliceSet[T] {
	var o options[T]
	for _, opt := range opts {
		opt(&o)
	}

	var s SliceSet[T]
	if o.initialCapacity > 0 {
		s.values = make([]T, 0, o.initialCapacity)
	}

	return &s
}

type SliceSet[T comparable] struct {
	values []T
}

func (s *SliceSet[T]) Add(v T) {
	if s.Contains(v) {
		return
	}

	s.values = append(s.values, v)
}

func (s *SliceSet[T]) Clear() {
	s.values = s.values[:0]
}

func (s *SliceSet[T]) Contains(v T) bool {
	for _, sv := range s.values {
		if v == sv {
			return true
		}
	}

	return false
}

func (s *SliceSet[T]) Iter() iter.Iter[T] {
	return iter.FromSlice[T](s.values)
}

func (s *SliceSet[T]) Len() int {
	return len(s.values)
}

func (s *SliceSet[T]) Remove(v T) {
	for i, sv := range s.values {
		if v == sv {
			s.values[i] = s.values[len(s.values)-1]
			s.values = s.values[:len(s.values)]
			return
		}
	}
}
