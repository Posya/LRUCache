package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache[int](3)

	assert.Equal(t, "[0] ", c.toString(), "list is empty")

	c.Add("1", 1)
	c.Add("2", 2)
	c.Add("3", 3)
	assert.Equal(t, "[3]321 123", c.toString(), "add 1, 2 and 3")

	ok := c.Add("2", 2)
	assert.Equal(t, "[3]231 132", c.toString(), "touch 2")
	assert.Equal(t, false, ok, "item in a cache")

	c.Add("4", 4)
	assert.Equal(t, "[3]423 324", c.toString(), "add 4; 1 is go away")

	c.Add("1", 1)
	assert.Equal(t, "[3]142 241", c.toString(), "add 1 back; 3 is go away")

	v, ok := c.Get("1")
	assert.Equal(t, 1, v, "get 1")
	assert.True(t, ok, "item in a cache")

	_, ok = c.Get("10")
	assert.False(t, ok, "item absent")

	ok = c.Remove("4")
	assert.True(t, ok, "item removed")
	assert.Equal(t, "[2]12 21", c.toString(), "remove 4")

	ok = c.Remove("10")
	assert.False(t, ok, "item not found in the cache")
	assert.Equal(t, "[2]12 21", c.toString(), "cache is not changed")
}

func TestLRUCacheAddEmptyKey(t *testing.T) {
	c := NewLRUCache[int](1)
	assert.Panics(t, func() { c.Add("", 10) }, "panics on empty key")
}
