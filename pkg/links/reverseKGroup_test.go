package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	links := []*ListNode{
		NumToList(54321),
		NumToList(1),
		NumToList(654321),
	}
	k := []int{
		2,
		1,
		2,
	}
	result := []*ListNode{
		NumToList(53412),
		NumToList(1),
		NumToList(563412),
	}
	for i, r := range links {
		rr := ReverseKGroup(r, k[i])
		assert.Equal(t, rr.String(), result[i].String())
	}
}
