package iter

type Collector[T any] interface {
	Add(...T)
}

func Collect[T any](dst Collector[T], it Iter[T]) error {
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		dst.Add(v)
	}

	return it.Close()
}
