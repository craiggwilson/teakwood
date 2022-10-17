package list

import "github.com/craiggwilson/teakwood/util/iter"

type linkedNode[T any] struct {
	value T
	next  *linkedNode[T]
	prev  *linkedNode[T]
}

func NewLinked[T any](values ...T) *Linked[T] {
	var l Linked[T]
	l.root.next = &l.root
	l.root.prev = &l.root

	if len(values) > 0 {
		l.Add(values...)
	}

	return &l
}

type Linked[T any] struct {
	len  int
	root linkedNode[T]
}

func (l *Linked[T]) Add(values ...T) {
	for _, v := range values {
		n := &linkedNode[T]{value: v}
		l.insertAt(n, l.root.prev)
	}
}

func (l *Linked[T]) InsertAt(idx int, values ...T) {
	at := l.nodeAt(idx)
	for _, v := range values {
		n := &linkedNode[T]{value: v}
		l.insertAt(n, at)
		at = n
	}
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

func (it *linkedIter[T]) Err() error {
	return nil
}
