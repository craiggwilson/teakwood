package iter

type Iterer[T any] interface {
	Iter() Iter[T]
}

type ItererFunc[T any] func() Iter[T]

func (f ItererFunc[T]) Iter() Iter[T] {
	return f()
}

type Iter[T any] interface {
	Next() (T, bool)
	Close() error
}
