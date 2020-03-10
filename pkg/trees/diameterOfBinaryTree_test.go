package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	trees := []*treeNode{
		&treeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		&treeNode{
			Val: 1,
			Left: &treeNode{
				Val: 2,
				Left: &treeNode{
					Val: 4,
				},
				Right: &treeNode{
					Val: 5,
				},
			},
			Right: &treeNode{
				Val: 3,
			},
		},
	}
	result := []int{
		0,
		3,
	}
	for i, r := range trees {
		rr := DiameterOfBinaryTree(r)
		assert.Equal(t, rr, result[i])
	}
}
