package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPostTree(t *testing.T) {
	lists := [][2][]int{
		{
			{9, 3, 15, 20, 7},
			{9, 15, 7, 20, 3},
		},
		{
			nil,
			nil,
		},
	}
	result := []*TreeNode{
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
	for i, r := range lists {
		rr := BuildPostTree(r[0], r[1])
		assert.Equal(t, rr.String(), result[i].String())
	}
}
