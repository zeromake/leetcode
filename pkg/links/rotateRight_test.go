package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateRight(t *testing.T) {
	links := []*ListNode{
		NumToList(54321),
		NumToList(21),
		NumToList(1),
	}
	k := []int{
		2,
		0,
		2,
	}
	result := []*ListNode{
		NumToList(32154),
		NumToList(21),
		NumToList(1),
	}
	for i, r := range links {
		rr := RotateRight(r, k[i])
		assert.Equal(t, rr.String(), result[i].String())
	}
}
