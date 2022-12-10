package lrucache

import "fmt"

type LRUCache[T any] interface {
	Add(string, T) bool
	Get(string) (T, bool)
	Remove(string) bool
}

type lruCache[T any] struct {
	l    list[T]
	m    map[string]*node[T]
	size int
}

func NewLRUCache[T any](size int) *lruCache[T] {
	cache := lruCache[T]{size: size, m: make(map[string]*node[T])}
	return &cache
}

func (c *lruCache[T]) cutIfOversized() {
	if len(c.m) == c.size {
		nodeToDelete := c.l.tail
		delete(c.m, nodeToDelete.key)
		c.l.cut(nodeToDelete)
	}
}

func (c *lruCache[T]) Add(key string, value T) bool {
	if key == "" {
		panic("key is empty string")
	}
	if n, ok := c.m[key]; ok {
		c.l.cut(n)
		c.l.addHead(n)
		return false
	}
	newNode := node[T]{key: key, data: value}
	c.cutIfOversized()
	c.m[key] = &newNode
	c.l.addHead(&newNode)
	return true
}

func (c *lruCache[T]) Get(key string) (T, bool) {
	var value T
	n, ok := c.m[key]
	if !ok {
		return value, false
	}
	c.l.cut(n)
	c.l.addHead(n)
	return n.data, true
}

func (c *lruCache[T]) Remove(key string) bool {
	n, ok := c.m[key]
	if !ok {
		return false
	}
	delete(c.m, key)
	c.l.cut(n)
	return true
}

func (c *lruCache[T]) toString() string {
	return fmt.Sprintf("[%d]%s", len(c.m), c.l.toString())
}
