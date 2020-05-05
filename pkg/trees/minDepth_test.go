package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDepth(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 3,
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
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		nil,
	}
	result := []int{
		2,
		2,
		0,
	}
	for i, r := range trees {
		rr := MinDepth(r)
		assert.Equal(t, rr, result[i])
	}
}
