package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreOrderTraversal(t *testing.T) {
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
		nil,
	}
	result := [][]int{
		{1, 2, 3},
		nil,
	}
	for i, r := range trees {
		rr := PreOrderTraversal(r)
		assert.Equal(t, rr, result[i])
	}
}
