package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	loopStart := &ListNode{
		Val: -4,
	}
	loopEnd := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: loopStart,
		},
	}
	loopStart.Next = loopEnd
	links := []*ListNode{
		{
			Val:  3,
			Next: loopEnd,
		},
		NumToList(123456),
		nil,
	}
	result := []bool{
		true,
		false,
		false,
	}
	for i, r := range links {
		rr := HasCycle(r)
		assert.Equal(t, rr, result[i])
	}
}
