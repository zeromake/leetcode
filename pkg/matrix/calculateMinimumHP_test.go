package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	dungeon := [][][]int{
		{
			{-2, -3, 3},
			{-5, -10, 1},
			{10, 30, -5},
		},
	}
	result := []int{
		7,
	}
	for i, r := range dungeon {
		rr := CalculateMinimumHP(r)
		assert.Equal(t, rr, result[i])
	}
}
