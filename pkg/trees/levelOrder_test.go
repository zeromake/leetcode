package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	trees := []*TreeNode{
		{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		nil,
	}
	result := [][][]int{
		{
			{3},
			{9, 20},
			{15, 7},
		},
		nil,
	}
	for i, r := range trees {
		rr := LevelOrder(r)
		assert.Equal(t, rr, result[i])
	}
}
