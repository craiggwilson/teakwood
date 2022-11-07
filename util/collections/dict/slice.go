package dict

import (
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ Dict[int, int] = (*Slice[int, int])(nil)

func NewSlice[K comparable, V any]() *Slice[K, V] {
	return &Slice[K, V]{}
}

type Slice[K comparable, V any] struct {
	pairs []iter.KeyValuePair[K, V]
}

func (d *Slice[K, V]) Add(k K, v V) {
	i := d.indexOf(k)
	if i >= 0 {
		d.pairs[i] = iter.KeyValuePair[K, V]{k, v}
	} else {
		d.pairs = append(d.pairs, iter.KeyValuePair[K, V]{k, v})
	}
}

func (d *Slice[K, V]) Contains(k K) bool {
	i := d.indexOf(k)

	return i >= 0
}

func (d *Slice[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return iter.FromSlice[iter.KeyValuePair[K, V]](d.pairs)
}

func (d *Slice[K, V]) Keys() iter.Iter[K] {
	return iter.Project(d.Iter(), func(kvp iter.KeyValuePair[K, V]) K {
		return kvp.Key
	})
}

func (d *Slice[K, V]) Len() int {
	return len(d.pairs)
}

func (d *Slice[K, V]) Remove(k K) {
	i := d.indexOf(k)
	if i < 0 {
		return
	}

	d.pairs[i] = d.pairs[len(d.pairs)-1]
	d.pairs = d.pairs[:len(d.pairs)]
}

func (d *Slice[K, V]) Value(k K) (V, bool) {
	i := d.indexOf(k)
	if i < 0 {
		var def V
		return def, false
	}

	return d.pairs[i].Value, true
}

func (d *Slice[K, V]) Values() iter.Iter[V] {
	return iter.Project(d.Iter(), func(kvp iter.KeyValuePair[K, V]) V {
		return kvp.Value
	})
}

func (d *Slice[K, V]) indexOf(k K) int {
	for i, pair := range d.pairs {
		if pair.Key == k {
			return i
		}
	}

	return -1
}
