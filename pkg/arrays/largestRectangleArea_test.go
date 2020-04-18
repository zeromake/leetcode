package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	heights := [][]int{
		{2, 1, 5, 6, 2, 3},
		nil,
		{3},
		{1, 1},
	}
	result := []int{
		10,
		0,
		3,
		2,
	}
	for i, r := range heights {
		rr := LargestRectangleArea(r)
		assert.Equal(t, rr, result[i])
	}
}
