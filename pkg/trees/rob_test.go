package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	result := []int{
		7,
		9,
	}
	for i, r := range trees {
		rr := Rob(r)
		assert.Equal(t, rr, result[i])
	}
}
