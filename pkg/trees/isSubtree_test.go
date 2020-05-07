package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubtree(t *testing.T) {
	trees := [][2]*TreeNode{
		{
			{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 2,
					},
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
					},
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range trees {
		rr := IsSubtree(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
