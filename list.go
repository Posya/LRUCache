package lrucache

import (
	"strings"
)

type node[T any] struct {
	prev *node[T]
	next *node[T]
	key  string
	data T
}

type list[T any] struct {
	head *node[T]
	tail *node[T]
}

func (l *list[T]) addHead(newHead *node[T]) {
	if newHead == nil {
		panic("adding nil element")
	}
	newHead.next = l.head
	newHead.prev = nil
	if l.head != nil {
		l.head.prev = newHead
	}
	l.head = newHead
	if l.tail == nil {
		l.tail = newHead
	}
}

func (l *list[T]) link(n1, n2 *node[T]) {
	if n1 != nil {
		n1.next = n2
	}
	if n2 != nil {
		n2.prev = n1
	}

}

func (l *list[T]) cut(cutNode *node[T]) {
	if cutNode == nil {
		panic("cutting nil element")
	}
	prev := cutNode.prev
	next := cutNode.next
	l.link(prev, next)
	if l.head == cutNode {
		l.head = next
	}
	if l.tail == cutNode {
		l.tail = prev
	}
	cutNode.prev = nil
	cutNode.next = nil
}

func (l *list[T]) toString() string {
	var res strings.Builder
	for n := l.head; n != nil; n = n.next {
		res.WriteString(n.key)
	}
	res.WriteString(" ")
	for n := l.tail; n != nil; n = n.prev {
		res.WriteString(n.key)
	}

	return res.String()
}
