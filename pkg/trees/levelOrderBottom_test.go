package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
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
			{15, 7},
			{9, 20},
			{3},
		},
		nil,
	}
	for i, r := range trees {
		rr := LevelOrderBottom(r)
		assert.Equal(t, rr, result[i])
	}
}
