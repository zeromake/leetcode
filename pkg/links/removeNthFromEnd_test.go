package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	links := []*ListNode{
		NumToList(54321),
	}
	result := []*ListNode{
		NumToList(5321),
	}
	for i, r := range links {
		rr := RemoveNthFromEnd(r, 2)
		assert.Equal(t, rr.String(), result[i].String())
	}
}

