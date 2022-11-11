package iter

func ElementAt[T any](it Iter[T], idx int) (T, bool) {
	if idx < 0 {
		panic("idx must be greater than 0")
	}
	pos := 0
	for e, ok := it.Next(); ok; e, ok = it.Next() {
		if idx == pos {
			return e, true
		}
		pos++
	}

	var def T
	return def, false
}
