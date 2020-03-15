package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	arrs := [][]int{
		{
			-1, 2, 1, -4,
		},
		{
			1,
			2,
		},
		{
			-9, -1, 2, 1, -4, 8,
		},
		{
			0,
			0,
			0,
		},
	}
	targets := []int{
		1,
		2,
		2,
		1,
	}
	result := []int{
		2,
		-1,
		2,
		0,
	}
	for i, r := range arrs {
		rr := ThreeSumClosest(r, targets[i])
		assert.Equal(t, rr, result[i])
	}
}
