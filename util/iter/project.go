package iter

func Project[From, To any](it Iter[From], projector func(From) To) Iter[To] {
	return &projectIter[From, To]{
		it:        it,
		projector: projector,
	}
}

type projectIter[From, To any] struct {
	it        Iter[From]
	projector func(From) To
}

func (it *projectIter[From, To]) Next() (To, bool) {
	n, ok := it.it.Next()
	if !ok {
		var def To
		return def, false
	}

	return it.projector(n), true
}

func (it *projectIter[From, To]) Close() error {
	return it.it.Close()
}
