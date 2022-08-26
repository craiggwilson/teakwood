package items

type Source[T any] interface {
	Item(int) T
	Len() int
}
