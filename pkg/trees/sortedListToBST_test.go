package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	nums := []*ListNode{
		{
			Val: -10,
			Next: &ListNode{
				Val: -3,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
		},
	}
	result := []*TreeNode{
		{
			Val: 0,
			Left: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	for i, r := range nums {
		rr := SortedListToBST(r)
		assert.Equal(t, rr, result[i])
	}
}
