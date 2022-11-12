package mapdict

import (
	"github.com/craiggwilson/teakwood/util/collections/dict"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ dict.Dict[int, int] = (*MapDict[int, int])(nil)

func New[K comparable, V any](opts ...Opt[K, V]) *MapDict[K, V] {
	var o options[K, V]
	for _, opt := range opts {
		opt(&o)
	}

	var m map[K]V
	if o.initialCapacity > 0 {
		m = make(map[K]V, o.initialCapacity)
	} else {
		m = make(map[K]V)
	}

	return &MapDict[K, V]{values: m}
}

type MapDict[K comparable, V any] struct {
	values map[K]V
}

func (d *MapDict[K, V]) Add(k K, v V) {
	d.values[k] = v
}

func (d *MapDict[K, V]) Contains(k K) bool {
	_, ok := d.values[k]
	return ok
}

func (d *MapDict[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return iter.FromMap(d.values)
}

func (d *MapDict[K, V]) Keys() iter.Iter[K] {
	return iter.Map(
		iter.FromMap(d.values),
		func(kvp iter.KeyValuePair[K, V]) K {
			return kvp.Key
		},
	)
}

func (d *MapDict[K, V]) Len() int {
	return len(d.values)
}

func (d *MapDict[K, V]) Remove(k K) {
	delete(d.values, k)
}

func (d *MapDict[K, V]) Value(k K) (V, bool) {
	v, ok := d.values[k]
	if !ok {
		return v, false
	}

	return v, true
}

func (d *MapDict[K, V]) Values() iter.Iter[V] {
	return iter.Map(iter.FromMap(d.values), func(kvp iter.KeyValuePair[K, V]) V {
		return kvp.Value
	})
}
