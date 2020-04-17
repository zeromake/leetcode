package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
		},
		{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
		},
	}
	result := []bool{
		true,
		false,
		false,
		false,
	}
	for i, r := range trees {
		rr := IsValidBST(r)
		assert.Equal(t, rr, result[i])
	}
}
