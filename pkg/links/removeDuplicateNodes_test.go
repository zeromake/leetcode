package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	links := []*ListNode{
		NumToList(123321),
		NumToList(21111),
	}
	result := []*ListNode{
		NumToList(321),
		NumToList(21),
	}
	for i, r := range links {
		rr := RemoveDuplicateNodes(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}
