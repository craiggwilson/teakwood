package mapset

import (
	"github.com/craiggwilson/teakwood/util/collections/set"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ set.Set[int] = (*MapSet[int])(nil)

func New[T comparable](opts ...Opt[T]) *MapSet[T] {
	var o options[T]
	for _, opt := range opts {
		opt(&o)
	}

	var m map[T]struct{}
	if o.initialCapacity > 0 {
		m = make(map[T]struct{}, o.initialCapacity)
	} else {
		m = make(map[T]struct{})
	}

	return &MapSet[T]{m}
}

type MapSet[T comparable] struct {
	values map[T]struct{}
}

func (s *MapSet[T]) Add(v T) {
	s.values[v] = struct{}{}
}

func (s *MapSet[T]) Clear() {
	s.values = make(map[T]struct{})
}

func (s *MapSet[T]) Contains(v T) bool {
	_, ok := s.values[v]
	return ok
}

func (s *MapSet[T]) Iter() iter.Iter[T] {
	return iter.Project(
		iter.FromMap(s.values),
		func(kvp iter.KeyValuePair[T, struct{}]) T {
			return kvp.Key
		},
	)
}

func (s *MapSet[T]) Len() int {
	return len(s.values)
}

func (s *MapSet[T]) Remove(v T) {
	delete(s.values, v)
}
