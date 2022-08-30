package items

func NewSlice[T any](items ...T) *Slice[T] {
	return &Slice[T]{
		items: items,
	}
}

type Slice[T any] struct {
	items []T
}

func (m *Slice[T]) Add(items ...T) {
	m.items = append(m.items, items...)
}

func (m *Slice[T]) Clear() {
	m.items = m.items[:0]
}

func (m *Slice[T]) InsertAt(item T, index int) {
	m.items = append(m.items[:index+1], m.items[index:]...)
	m.items[index] = item
}

func (s *Slice[T]) Item(index int) T {
	return s.items[index]
}

func (s *Slice[T]) Len() int {
	return len(s.items)
}

func (m *Slice[T]) RemoveAt(index int) {
	m.items = append(m.items[:index], m.items[index+1:]...)
}
