package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	trees := [][2]*TreeNode{
		{
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	result := []bool{
		true,
		false,
		false,
	}
	for i, r := range trees {
		rr := IsSameTree(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
