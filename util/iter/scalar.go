package iter

import (
	"golang.org/x/exp/constraints"
)

func All[S any](src Iterer[S], predicate func(S) bool) (result bool, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	result = true
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if !predicate(elem) {
			result = false
			break
		}
	}

	return
}

func Any[S any](src Iterer[S], predicate func(S) bool) (result bool, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if predicate(elem) {
			result = true
			break
		}
	}

	return
}

func Collect[S any](src Iterer[S], dst interface{ Add(S) }) (err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		dst.Add(elem)
	}

	return
}

func Contains[S comparable](src Iterer[S], target S) (result bool, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if elem == target {
			result = true
			break
		}
	}

	return
}

func ContainsBy[S comparable](src Iterer[S], predicate func(S) bool) (result bool, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if predicate(elem) {
			result = true
			break
		}
	}

	return
}

func ElementAt[S any](src Iterer[S], idx uint) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	pos := uint(0)
	found := false
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if idx == pos {
			result = elem
			found = true
			break
		}
		pos++
	}

	if !found {
		err = makeOutOfRangeError("idx")
	}

	return
}

func First[S any](src Iterer[S]) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	elem, ok := it.Next()
	if !ok {
		err = makeEmptyError("src")
	} else {
		result = elem
	}

	return
}

func Fold[S, R any](src Iterer[S], seed R, reducer func(R, S) R) (result R, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	result = seed
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		result = reducer(result, elem)
	}

	return
}

func Last[S any](src Iterer[S]) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	isEmpty := true
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		result = elem
		isEmpty = false
	}

	if isEmpty {
		err = makeEmptyError("src")
	}

	return
}

func Len[S any, R constraints.Integer](src Iterer[S]) (result R, err error) {
	if counter, ok := src.(interface{ Len() R }); ok {
		result, err = counter.Len(), nil
	} else {
		it := src.Iter()
		defer func() {
			err = it.Close()
		}()

		for _, ok = it.Next(); ok; _, ok = it.Next() {
			result++
		}
	}

	return
}

func Max[S constraints.Ordered](src Iterer[S]) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	isEmpty := true
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if isEmpty {
			result = elem
			isEmpty = false
			continue
		}

		if elem > result {
			result = elem
		}
	}

	if isEmpty {
		err = makeEmptyError("src")
	}

	return
}

func Min[S constraints.Ordered](src Iterer[S]) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	isEmpty := true
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		if isEmpty {
			result = elem
			isEmpty = false
			continue
		}

		if elem < result {
			result = elem
		}
	}

	if isEmpty {
		err = makeEmptyError("src")
	}

	return
}

func Reduce[S any](src Iterer[S], reducer func(S, S) S) (result S, err error) {
	it := src.Iter()
	defer func() {
		cerr := it.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	elem, ok := it.Next()
	if !ok {
		err = makeEmptyError("src")
	} else {
		result = elem
		for elem, ok = it.Next(); ok; elem, ok = it.Next() {
			result = reducer(result, elem)
		}
	}

	return
}

func Sum[S constraints.Integer | constraints.Float](src Iterer[S]) (result S, err error) {
	it := src.Iter()
	defer func() {
		err = it.Close()
	}()

	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		result += elem
	}

	return
}

func ToSlice[S any](src Iterer[S]) (result []S, err error) {
	if slicer, ok := src.(interface{ ToSlice() []S }); ok {
		result, err = slicer.ToSlice(), nil
	} else {
		it := src.Iter()
		defer func() {
			err = it.Close()
		}()

		for elem, ok := it.Next(); ok; elem, ok = it.Next() {
			result = append(result, elem)
		}
	}

	return
}
