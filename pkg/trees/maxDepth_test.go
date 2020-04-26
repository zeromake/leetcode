package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	trees := []*TreeNode{
		{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		nil,
	}
	result := []int{
		3,
		0,
	}
	for i, r := range trees {
		rr := MaxDepth(r)
		assert.Equal(t, rr, result[i])
	}
}
