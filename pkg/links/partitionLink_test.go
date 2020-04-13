package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionLink(t *testing.T) {
	links := []*ListNode{
		NumToList(252341),
	}
	x := []int{
		3,
	}
	result := []*ListNode{
		NumToList(534221),
	}
	for i, r := range links {
		rr := PartitionLink(r, x[i])
		assert.Equal(t, rr, result[i])
	}
}
