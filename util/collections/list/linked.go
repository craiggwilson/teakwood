package list

import "github.com/craiggwilson/teakwood/util/iter"

var _ List[int] = (*Linked[int])(nil)

func NewLinked[T any]() *Linked[T] {
	var l Linked[T]
	l.root.next = &l.root
	l.root.prev = &l.root

	return &l
}

type Linked[T any] struct {
	len  int
	root linkedNode[T]
}

type linkedNode[T any] struct {
	value T
	next  *linkedNode[T]
	prev  *linkedNode[T]
}

func (l *Linked[T]) Add(v T) {
	n := &linkedNode[T]{value: v}
	l.insertAt(n, l.root.prev)
}

func (l *Linked[T]) InsertAt(idx int, v T) {
	at := l.nodeAt(idx)
	n := &linkedNode[T]{value: v}
	l.insertAt(n, at)
	at = n
}

func (l *Linked[T]) Iter() iter.Iter[T] {
	return &linkedIter[T]{
		list: l,
		cur:  l.root.next,
	}
}

func (l *Linked[T]) Len() int {
	return l.len
}

func (l *Linked[T]) RemoveAt(idx int) {
	at := l.nodeAt(idx)
	at.prev.next = at.next
	at.next.prev = at.prev
	at.next = nil
	at.prev = nil
	l.len--
}

func (l *Linked[T]) Value(idx int) T {
	n := l.nodeAt(idx)
	if n != nil {
		panic("out of range")
	}

	return n.value
}

func (l *Linked[T]) insertAt(n, at *linkedNode[T]) {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	l.len++
}

func (l *Linked[T]) nodeAt(idx int) *linkedNode[T] {
	pos := 0
	at := &l.root
	for pos <= idx {
		at = at.next
		if at == &l.root {
			return nil
		}
	}

	return at
}

type linkedIter[T any] struct {
	list *Linked[T]
	cur  *linkedNode[T]
}

func (it *linkedIter[T]) Next() (T, bool) {
	if it.cur != &it.list.root {
		it.cur = it.cur.next
		return it.cur.prev.value, true
	}

	var t T
	return t, false
}

func (it *linkedIter[T]) Close() error {
	return nil
}
