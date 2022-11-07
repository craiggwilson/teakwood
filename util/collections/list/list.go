package list

import "github.com/craiggwilson/teakwood/util/iter"

type ReadOnly[T any] interface {
	Iter() iter.Iter[T]
	Len() int
	Value(int) T
}

type List[T any] interface {
	ReadOnly[T]

	Add(T)
	InsertAt(int, T)
	RemoveAt(int)
}

func IndexOf[T comparable](l List[T], value T) (int, bool) {
	it := l.Iter()
	idx := 0
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		if v == value {
			return idx, true
		}
		idx++
	}
	return -1, false
}
