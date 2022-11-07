package set

import "github.com/craiggwilson/teakwood/util/iter"

type ReadOnly[T comparable] interface {
	Contains(T) bool
	Len() int
	Iter() iter.Iter[T]
}

type Set[T comparable] interface {
	ReadOnly[T]

	Add(T)
	Clear()
	Remove(T)
}
