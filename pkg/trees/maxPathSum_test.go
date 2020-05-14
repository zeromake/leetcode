package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	trees := []*TreeNode{
		{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		{
			Val: -10,
			Left: &TreeNode{
				Val: 9,
			},
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
	}
	result := []int{
		6,
		42,
	}
	for i, r := range trees {
		rr := MaxPathSum(r)
		assert.Equal(t, rr, result[i])
	}
}
