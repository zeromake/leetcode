package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	links := [][2]*ListNode{
		{
			NumToList(7531),
			NumToList(8642),
		},
		{
			NumToList(8642),
			NumToList(7531),
		},
	}
	result := []*ListNode{
		NumToList(87654321),
		NumToList(87654321),
	}
	for i, r := range links {
		rr := MergeTwoLists(r[0], r[1])
		assert.Equal(t, rr.String(), result[i].String())
	}
}
