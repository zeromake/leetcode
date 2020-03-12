package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	arrs := [][]int{
		{
			-1,
			0,
			0,
			1,
			1,
			2,
			-1,
			-1,
			-4,
		},
		{
			1,
			2,
		},
	}
	result := [][][]int{
		{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		{
		},
	}

	for i, r := range arrs {
		rr := ThreeSum(r)
		assert.Equal(t, rr, result[i])

		rr = ThreeSumTarget(r, 0)
		assert.Equal(t, rr, result[i])
	}
}
