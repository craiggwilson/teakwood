package slicedict

import (
	"github.com/craiggwilson/teakwood/util/collections/dict"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ dict.Dict[int, int] = (*SliceDict[int, int])(nil)

func New[K comparable, V any](opts ...Opt[K, V]) *SliceDict[K, V] {
	var o options[K, V]
	for _, opt := range opts {
		opt(&o)
	}

	var s SliceDict[K, V]
	if o.initialCapacity > 0 {
		s.pairs = make([]iter.KeyValuePair[K, V], 0, o.initialCapacity)
	}

	return &SliceDict[K, V]{}
}

type SliceDict[K comparable, V any] struct {
	pairs []iter.KeyValuePair[K, V]
}

func (d *SliceDict[K, V]) Add(k K, v V) {
	i := d.indexOf(k)
	if i >= 0 {
		d.pairs[i] = iter.KeyValuePair[K, V]{k, v}
	} else {
		d.pairs = append(d.pairs, iter.KeyValuePair[K, V]{k, v})
	}
}

func (d *SliceDict[K, V]) Contains(k K) bool {
	i := d.indexOf(k)

	return i >= 0
}

func (d *SliceDict[K, V]) Iter() iter.Iter[iter.KeyValuePair[K, V]] {
	return iter.FromSlice[iter.KeyValuePair[K, V]](d.pairs)
}

func (d *SliceDict[K, V]) Keys() iter.Iter[K] {
	return iter.Map(d.Iter(), func(kvp iter.KeyValuePair[K, V]) K {
		return kvp.Key
	})
}

func (d *SliceDict[K, V]) Len() int {
	return len(d.pairs)
}

func (d *SliceDict[K, V]) Remove(k K) {
	i := d.indexOf(k)
	if i < 0 {
		return
	}

	d.pairs[i] = d.pairs[len(d.pairs)-1]
	d.pairs = d.pairs[:len(d.pairs)]
}

func (d *SliceDict[K, V]) Value(k K) (V, bool) {
	i := d.indexOf(k)
	if i < 0 {
		var def V
		return def, false
	}

	return d.pairs[i].Value, true
}

func (d *SliceDict[K, V]) Values() iter.Iter[V] {
	return iter.Map(d.Iter(), func(kvp iter.KeyValuePair[K, V]) V {
		return kvp.Value
	})
}

func (d *SliceDict[K, V]) indexOf(k K) int {
	for i, pair := range d.pairs {
		if pair.Key == k {
			return i
		}
	}

	return -1
}
