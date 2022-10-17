package list

import "github.com/craiggwilson/teakwood/util/iter"

func NewFrozen[T any](l ReadOnly[T]) *Frozen[T] {
	return &Frozen[T]{l}
}

type Frozen[T any] struct {
	l ReadOnly[T]
}

func (l *Frozen[T]) Iter() iter.Iter[T] {
	return l.l.Iter()
}

func (l *Frozen[T]) Len() int {
	return l.l.Len()
}

func (l *Frozen[T]) Value(idx int) T {
	return l.l.Value(idx)
}
