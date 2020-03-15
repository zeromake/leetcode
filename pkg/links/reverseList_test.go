package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	links := []*ListNode{
		NumToList(54321),
	}
	result := []*ListNode{
		NumToList(12345),
	}
	for i, r := range links {
		rr := ReverseList(r)
		assert.Equal(t, rr, result[i])
	}
}
