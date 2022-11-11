package dlinkedlist

import "github.com/craiggwilson/teakwood/util/collections/list"
import "github.com/craiggwilson/teakwood/util/iter"

var _ list.List[int] = (*DLinkedList[int])(nil)

func New[T any]() *DLinkedList[T] {
	var l DLinkedList[T]
	l.root.next = &l.root
	l.root.prev = &l.root

	return &l
}

type DLinkedList[T any] struct {
	len  int
	root node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func (l *DLinkedList[T]) Add(v T) {
	n := &node[T]{value: v}
	l.insertAt(n, l.root.prev)
}

func (l *DLinkedList[T]) InsertAt(idx int, v T) {
	at := l.nodeAt(idx)
	n := &node[T]{value: v}
	l.insertAt(n, at)
	at = n
}

func (l *DLinkedList[T]) Iter() iter.Iter[T] {
	return &linkedIter[T]{
		list: l,
		cur:  l.root.next,
	}
}

func (l *DLinkedList[T]) Len() int {
	return l.len
}

func (l *DLinkedList[T]) RemoveAt(idx int) {
	at := l.nodeAt(idx)
	at.prev.next = at.next
	at.next.prev = at.prev
	at.next = nil
	at.prev = nil
	l.len--
}

func (l *DLinkedList[T]) Value(idx int) T {
	n := l.nodeAt(idx)
	if n != nil {
		panic("out of range")
	}

	return n.value
}

func (l *DLinkedList[T]) insertAt(n, at *node[T]) {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	l.len++
}

func (l *DLinkedList[T]) nodeAt(idx int) *node[T] {
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
	list *DLinkedList[T]
	cur  *node[T]
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
