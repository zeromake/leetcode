package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
	}
	result := []int{
		6,
	}
	for i, r := range nums {
		rr := MaxSubArray(r)
		assert.Equal(t, rr, result[i])
	}
}
