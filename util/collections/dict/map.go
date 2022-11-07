package dict

import "github.com/craiggwilson/teakwood/util/iter"

var _ Dict[int, int] = (*Map[int, int])(nil)

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{values: make(map[K]V)}
}

type Map[K comparable, V any] struct {
	values map[K]V
}

func (d *Map[K, V]) Add(k K, v V) {
	d.values[k] = v
}

func (d *Map[K, V]) Contains(k K) bool {
	_, ok := d.values[k]
	return ok
}

func (d *Map[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return iter.FromMap(d.values)
}

func (d *Map[K, V]) Keys() iter.Iter[K] {
	return iter.Project(iter.FromMap(d.values), func(kvp iter.KeyValuePair[K, V]) K {
		return kvp.Key
	})
}

func (d *Map[K, V]) Len() int {
	return len(d.values)
}

func (d *Map[K, V]) Remove(k K) {
	delete(d.values, k)
}

func (d *Map[K, V]) Value(k K) (V, bool) {
	v, ok := d.values[k]
	if !ok {
		return v, false
	}

	return v, true
}

func (d *Map[K, V]) Values() iter.Iter[V] {
	return iter.Project(iter.FromMap(d.values), func(kvp iter.KeyValuePair[K, V]) V {
		return kvp.Value
	})
}
