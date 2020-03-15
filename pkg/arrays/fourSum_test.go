package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFourSum(t *testing.T) {
	arrs := [][]int{
		{
			1,
			9,
			9,
			0,
			-1,
			0,
			-2,
			2,
			2,
		},
		{
			0,
			0,
			0,
			0,
		},
		{},
	}
	result := [][][]int{
		{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		},
		{
			{0, 0, 0, 0},
		},
		{},
	}
	for i, r := range arrs {
		rr := FourSum(r, 0)
		assert.Equal(t, rr, result[i])
	}
}
