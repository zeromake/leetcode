package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRightSideView(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		nil,
	}
	result := [][]int{
		{1, 3, 4},
		{1, 3, 4},
		nil,
	}
	for i, r := range trees {
		rr := RightSideView(r)
		assert.Equal(t, rr, result[i])
	}
}
