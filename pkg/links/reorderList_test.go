package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReorderList(t *testing.T) {
	links := []*ListNode{
		NumToList(4321),
		NumToList(54321),
	}
	result := []*ListNode{
		NumToList(3241),
		NumToList(34251),
	}
	for i, r := range links {
		ReorderList(r)
		assert.Equal(t, r.String(), result[i].String())
	}
}
