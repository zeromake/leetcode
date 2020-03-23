package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNode(t *testing.T)  {
	links := []*ListNode{
		NumToList(54321),
		NumToList(654321),
	}
	result := []*ListNode{
		NumToList(543),
		NumToList(654),
	}
	for i, r := range links {
		rr := MiddleNode(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}
