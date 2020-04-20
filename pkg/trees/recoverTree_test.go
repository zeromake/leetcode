package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			Val: 0, Left: &TreeNode{Val: 1},
		},
	}
	result := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			Val: 1, Left: &TreeNode{Val: 0},
		},
	}
	for i, r := range trees {
		RecoverTree(r)
		assert.Equal(t, r.String(), result[i].String())
	}
}
