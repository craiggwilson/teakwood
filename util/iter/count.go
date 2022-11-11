package iter

import "golang.org/x/exp/constraints"

func Count[T any, R constraints.Integer](it Iter[T]) R {
	var r R
	for _, ok := it.Next(); ok; _, ok = it.Next() {
		r++
	}
	return r
}
