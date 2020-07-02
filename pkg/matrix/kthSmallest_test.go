package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	matrix := [][][]int{
		{
			{1, 5, 9},
			{10, 11, 13},
			{12, 13, 15},
		},
		{
			{1, 2},
			{1, 3},
		},
		{
			{-5},
		},
	}
	k := []int{
		8,
		2,
		1,
	}
	result := []int{
		13,
		1,
		-5,
	}
	for i, r := range matrix {
		rr := KthSmallest(r, k[i])
		assert.Equal(t, rr, result[i])
	}
}
