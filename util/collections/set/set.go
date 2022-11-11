package set

import "github.com/craiggwilson/teakwood/util/iter"

type ReadOnly[T comparable] interface {
	iter.Iterer[T]

	Contains(T) bool
	Len() int
}

type Set[T comparable] interface {
	ReadOnly[T]

	Add(T)
	Clear()
	Remove(T)
}

func AddFromIter[T comparable](s Set[T], it iter.Iter[T]) {
	for e, ok := it.Next(); ok; e, ok = it.Next() {
		s.Add(e)
	}
}

func AddFromSlice[T comparable](s Set[T], slice []T) {
	for _, e := range slice {
		s.Add(e)
	}
}
