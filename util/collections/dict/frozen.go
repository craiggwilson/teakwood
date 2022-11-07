package dict

import "github.com/craiggwilson/teakwood/util/iter"

var _ ReadOnly[int, int] = (*Frozen[int, int])(nil)

func NewFrozen[K comparable, V any](l ReadOnly[K, V]) *Frozen[K, V] {
	return &Frozen[K, V]{l}
}

type Frozen[K comparable, V any] struct {
	d ReadOnly[K, V]
}

func (d *Frozen[K, V]) Contains(k K) bool {
	return d.d.Contains(k)
}

func (d *Frozen[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return d.d.Iter()
}

func (d *Frozen[K, V]) Keys() iter.Iter[K] {
	return d.d.Keys()
}

func (d *Frozen[K, V]) Len() int {
	return d.d.Len()
}

func (d *Frozen[K, V]) Value(k K) (V, bool) {
	return d.d.Value(k)
}

func (d *Frozen[K, V]) Values() iter.Iter[V] {
	return d.d.Values()
}
