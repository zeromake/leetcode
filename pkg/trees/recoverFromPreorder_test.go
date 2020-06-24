package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecoverFromPreorder(t *testing.T) {
	strs := []string{
		"1-2--3--4-5--6--7",
		"1-2--3---4-5--6---7",
		"1-401--349---90--88",
		"",
	}
	result := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
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
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val: 401,
				Left: &TreeNode{
					Val: 349,
					Left: &TreeNode{
						Val: 90,
					},
				},
				Right: &TreeNode{
					Val: 88,
				},
			},
		},
		nil,
	}
	for i, r := range strs {
		rr := RecoverFromPreorder(r)
		assert.Equal(t, rr.String(), result[i].String())
	}
}
