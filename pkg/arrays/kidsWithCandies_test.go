package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	candies := [][]int{
		{2, 3, 5, 1, 3},
		{4, 2, 1, 1, 2},
		{12, 1, 12},
	}
	extraCandies := []int{
		3,
		1,
		10,
	}
	result := [][]bool{
		{true, true, true, false, true},
		{true, false, false, false, false},
		{true, false, true},
	}
	for i, r := range candies {
		rr := KidsWithCandies(r, extraCandies[i])
		assert.Equal(t, rr, result[i])
	}
}
