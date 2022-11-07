package dict

import (
	"github.com/craiggwilson/teakwood/util/iter"
)

type ReadOnly[K comparable, V any] interface {
	Contains(K) bool
	Iter() iter.Iter[iter.KeyValuePair[K, V]]
	Keys() iter.Iter[K]
	Len() int
	Value(K) (V, bool)
	Values() iter.Iter[V]
}

type Dict[K comparable, V any] interface {
	ReadOnly[K, V]

	Add(K, V)
	Remove(K)
}
