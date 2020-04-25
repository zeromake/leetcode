package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
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
		nil,
	}
	result := [][][]int{
		{
			{3},
			{20, 9},
			{15, 7},
		},
		nil,
	}
	for i, r := range trees {
		rr := ZigzagLevelOrder(r)
		assert.Equal(t, rr, result[i])
		rr = ZigzagLevelOrder2(r)
		assert.Equal(t, rr, result[i])
	}
}
