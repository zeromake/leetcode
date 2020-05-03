package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	trees := []*TreeNode{
		{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{Val: 2},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range trees {
		rr := IsBalanced(r)
		assert.Equal(t, rr, result[i])
	}
}
