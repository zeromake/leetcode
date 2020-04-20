package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	trees := []*TreeNode{
		{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 2},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
	}
	result := []bool{
		true,
		true,
		false,
	}
	for i, r := range trees {
		rr := IsSymmetric(r)
		assert.Equal(t, rr, result[i])
	}
}
