package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	links := []*ListNode{
		NumToList(54321),
		NumToList(53),
		NumToList(5),
	}
	n := [][2]int{
		{2, 4},
		{1, 1},
		{1, 1},
	}
	result := []*ListNode{
		NumToList(52341),
		NumToList(53),
		NumToList(5),
	}
	for i, r := range links {
		m := n[i]
		rr := ReverseBetween(r, m[0], m[1])
		assert.Equal(t, rr.String(), result[i].String())
	}
}
