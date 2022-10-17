package list

import "github.com/craiggwilson/teakwood/util/iter"

type ReadOnly[T any] interface {
	Iter() iter.Iter[T]
	Len() int
	Value(int) T
}

type List[T any] interface {
	ReadOnly[T]

	Add(...T)
	InsertAt(int, ...T)
	RemoveAt(int)
}
