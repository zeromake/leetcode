package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	n := []int{
		3,
		0,
	}
	result := []TreeNodeList{
		{
			{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
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
					Val: 3,
				},
			},
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
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		nil,
	}
	for i, r := range n {
		rr := GenerateTrees(r)
		assert.Equal(t, TreeNodeList(rr).String(), result[i].String())
	}
}

func TestTreeNode_String(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	result := []string{
		"1,null,3,2,null",
		"3,2,1,null,null",
	}
	for i, r := range trees {
		rr := r.String()
		assert.Equal(t, rr, result[i])
	}
}
