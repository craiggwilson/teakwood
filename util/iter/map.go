package iter

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

func FromMap[K comparable, V any](m map[K]V) Iter[KeyValuePair[K, V]] {
	entries := make([]KeyValuePair[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, KeyValuePair[K, V]{k, v})
	}

	return FromSlice(entries)
}
