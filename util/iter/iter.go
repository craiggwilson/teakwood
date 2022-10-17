package iter

type Iter[T any] interface {
	Next() (T, bool)
	Err() error
}
