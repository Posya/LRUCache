package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	l := list[int]{}

	n1 := node[int]{key: "1", data: 1}
	l.addHead(&n1)
	n2 := node[int]{key: "2", data: 2}
	l.addHead(&n2)

	assert.Equal(t, "21 12", l.toString(), "add 1 and 2")

	n3 := node[int]{key: "3", data: 3}
	l.addHead(&n3)

	assert.Equal(t, "321 123", l.toString(), "add 3")

	n4 := node[int]{key: "4", data: 4}
	l.addHead(&n4)
	n5 := node[int]{key: "5", data: 5}
	l.addHead(&n5)

	assert.Equal(t, "54321 12345", l.toString(), "add 4 and 5")

	l.cut(&n4)

	assert.Equal(t, "5321 1235", l.toString(), "cut 4")

	l.addHead(&n4)

	assert.Equal(t, "45321 12354", l.toString(), "add 4 back")

	l.cut(&n1)
	l.cut(&n2)
	l.cut(&n3)

	assert.Equal(t, "45 54", l.toString(), "cut 1, 2 and 3")

	l.cut(l.tail)

	assert.Equal(t, "4 4", l.toString(), "cut tail")

	l.cut(&n4)

	assert.Equal(t, " ", l.toString(), "cut last node")
}

func TestLinkedListAddNill(t *testing.T) {
	l := list[int]{}
	assert.Panics(t, func() { l.addHead(nil) }, "panics if nil in addHead")
	assert.Panics(t, func() { l.cut(nil) }, "panics if nil in cut")
}
