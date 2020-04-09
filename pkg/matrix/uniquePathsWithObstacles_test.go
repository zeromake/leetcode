package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	grid := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 0, 1},
			{1, 0, 0},
			{0, 0, 0},
		},
		{},
		{{}},
	}
	result := []int{
		2,
		2,
		0,
		0,
	}
	for i, r := range grid {
		rr := UniquePathsWithObstacles(r)
		assert.Equal(t, rr, result[i])
	}
}
