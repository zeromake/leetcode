package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	links := []*ListNode{
		NumToList(654321),
		NumToList(54321),
	}
	result := []*ListNode{
		NumToList(563412),
		NumToList(53412),
	}
	for i, r := range links {
		rr := SwapPairs(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}

