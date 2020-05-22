package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree_Rank(t *testing.T) {
	tree := &BinaryTree{
		value: 8,
		left: &BinaryTree{
			value: 4,
			left: &BinaryTree{
				value: 2,
			},
			right: &BinaryTree{
				value: 6,
			},
		},
		right: &BinaryTree{
			value: 12,
			left: &BinaryTree{
				value: 10,
			},
			right: &BinaryTree{
				value: 14,
			},
		},
	}
	tree1 := &BinaryTree{
	value: 2,
		left:  &BinaryTree{value: 1},
			right: &BinaryTree{
			value: 33,
			left: &BinaryTree{
				value: 25,
				left: &BinaryTree{
					value: 11,
					right: &BinaryTree{
						value: 7,
						right: &BinaryTree{
							value: 12,
							right: &BinaryTree{
								value: 13,
							},
						},
					},
				},
			},
			right: &BinaryTree{
				value: 40,
				left: &BinaryTree{
					value: 34,
					right: &BinaryTree{
						value: 36,
					},
				},
			},
		},
	}
	trees := []*BinaryTree{
		tree,
		tree,
		tree,
		tree,
		tree1,
		tree1,
	}
	num := []int{
		3,
		4,
		15,
		13,
		6,
		35,
	}
	result := []int{
		2,
		2,
		8,
		7,
		3,
		10,
	}
	for i, r := range trees {
		rr := r.Rank(num[i])
		assert.Equal(t, rr, result[i], i)
	}
}
