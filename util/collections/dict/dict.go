package dict

import (
	"github.com/craiggwilson/teakwood/util/iter"
)

type ReadOnly[K comparable, V any] interface {
	iter.Iterer[iter.KeyValuePair[K, V]]

	Contains(K) bool
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

func AddFromIter[K comparable, V any](d Dict[K, V], it iter.Iter[iter.KeyValuePair[K, V]]) {
	for e, ok := it.Next(); ok; e, ok = it.Next() {
		d.Add(e.Key, e.Value)
	}
}

func AddFromSlice[K comparable, V any](d Dict[K, V], slice []iter.KeyValuePair[K, V]) {
	for _, e := range slice {
		d.Add(e.Key, e.Value)
	}
}

func AddFromMap[K comparable, V any](d Dict[K, V], m map[K]V) {
	for k, v := range m {
		d.Add(k, v)
	}
}
