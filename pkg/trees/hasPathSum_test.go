package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: -2,
			Right: &TreeNode{
				Val: -3,
			},
		},
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
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			Val: -2,
			Right: &TreeNode{
				Val: -3,
			},
		},
		nil,
	}
	sum := []int{
		-5,
		22,
		-6,
		1,
	}
	result := []bool{
		true,
		true,
		false,
		false,
	}
	for i, r := range trees {
		rr := HasPathSum(r, sum[i])
		assert.Equal(t, rr, result[i])
	}
}
