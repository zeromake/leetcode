package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := [][]int{
		{-10, -3, 0, 5, 9},
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
		rr := SortedArrayToBST(r)
		assert.Equal(t, rr, result[i])
	}
}
