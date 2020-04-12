package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicatesII(t *testing.T) {
	links := []*ListNode{
		NumToList(5443321),
	}
	result := []*ListNode{
		NumToList(521),
	}
	for i, r := range links {
		rr := DeleteDuplicatesII(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}

func TestDeleteDuplicatesI(t *testing.T) {
	links := []*ListNode{
		NumToList(5443321),
	}
	result := []*ListNode{
		NumToList(54321),
	}
	for i, r := range links {
		rr := DeleteDuplicatesI(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}
