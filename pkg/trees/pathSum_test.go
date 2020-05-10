package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathSum(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	sum := []int{
		22,
	}
	result := [][][]int{
		{
			{5, 4, 11, 2},
			{5, 8, 4, 5},
		},
	}
	for i, r := range trees {
		rr := PathSum(r, sum[i])
		assert.Equal(t, rr, result[i])
	}
}
