package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	T := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
	}
	result := [][]int{
		{1, 1, 4, 2, 1, 1, 0, 0},
	}
	for i, r := range T {
		rr := DailyTemperatures(r)
		assert.Equal(t, rr, result[i])
	}
}
