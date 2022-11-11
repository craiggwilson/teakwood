package iter

func Reduce[T, R any](it Iter[T], initValue R, reducer func(R, T) R) R {
	cur := initValue
	for e, ok := it.Next(); ok; e, ok = it.Next() {
		cur = reducer(cur, e)
	}
	return cur
}
