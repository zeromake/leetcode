package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	one := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	tow := &ListNode{
		Val: -4,
		Next: one,
	}
	one.Next.Next = tow
	links := []*ListNode{
		{
			Val: 3,
			Next: one,
		},
		nil,
	}
	result := []*ListNode{
		{
			Val: 2,
			Next: nil,
		},
		nil,
	}
	for i, r := range links {
		rr := DetectCycle(r)
		if rr != nil && result[i] != nil {
			assert.Equal(t, rr.Val, result[i].Val)
			continue
		}
		if result[i] == nil {
			assert.Nil(t, rr)
		} else {
			assert.NotNil(t, rr)
		}
	}
}
