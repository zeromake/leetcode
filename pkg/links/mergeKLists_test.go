package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	links := [][]*ListNode{
		{
			NumToList(7531),
			NumToList(8642),
			NumToList(9),
		},
		{
			NumToList(9),
			NumToList(8642),
			NumToList(7531),
		},
		{
			NumToList(9),
			NumToList(8642),
			NumToList(753),
			NumToList(1),
		},
		{},
	}

	result := []*ListNode{
		NumToList(987654321),
		NumToList(987654321),
		NumToList(987654321),
		nil,
	}

	for i, r := range links {
		rr := MergeKLists(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}
