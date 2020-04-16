package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
	}
	result := [][]int{
		{1, 3, 2},
	}
	for i, r := range trees {
		rr := InorderTraversal(r)
		assert.Equal(t, rr, result[i])
	}
}
