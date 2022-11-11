package iter

func Zip[T1, T2, R any](it1 Iter[T1], it2 Iter[T2], zipper func(T1, T2) R) Iter[R] {
	return &zipIterator[T1, T2, R]{it1, it2, zipper}
}

type zipIterator[T1, T2, R any] struct {
	it1 Iter[T1]
	it2 Iter[T2]

	zipper func(T1, T2) R
}

func (it *zipIterator[T1, T2, R]) Next() (R, bool) {
	one, ok1 := it.it1.Next()
	two, ok2 := it.it2.Next()

	if ok1 && ok2 {
		return it.zipper(one, two), true
	}

	var def R
	return def, false
}

func (it *zipIterator[T1, T2, R]) Close() error {
	err1 := it.it1.Close()
	err2 := it.it2.Close()

	if err1 != nil {
		return err1
	}

	return err2
}
