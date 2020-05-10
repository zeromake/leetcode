package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		nil,
	}
	nodes := [][2]*TreeNode{
		{
			&TreeNode{
				Val: 5,
			},
			&TreeNode{
				Val: 1,
			},
		},
		{
			&TreeNode{
				Val: 5,
			},
			&TreeNode{
				Val: 4,
			},
		},
		{
			nil,
			nil,
		},
	}
	result := []*TreeNode{
		trees[0],
		trees[1].Left,
		nil,
	}
	for i, r := range trees {
		rr := LowestCommonAncestor(r, nodes[i][0], nodes[i][1])
		assert.Equal(t, rr, result[i])
		rr = LowestCommonAncestor2(r, nodes[i][0], nodes[i][1])
		assert.Equal(t, rr, result[i])
	}
}

func TestLowestCommonAncestor2(t *testing.T) {
	trees := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		nil,
	}
	nodes := [][2]*TreeNode{
		{
			&TreeNode{
				Val: 5,
			},
			&TreeNode{
				Val: 1,
			},
		},
		{
			&TreeNode{
				Val: 5,
			},
			&TreeNode{
				Val: 4,
			},
		},
		{
			nil,
			nil,
		},
	}
	result := []*TreeNode{
		trees[0],
		trees[1].Left,
		nil,
	}
	for i, r := range trees {
		rr := LowestCommonAncestor2(r, nodes[i][0], nodes[i][1])
		assert.Equal(t, rr, result[i])
	}
}
