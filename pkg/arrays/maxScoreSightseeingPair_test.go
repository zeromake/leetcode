package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScoreSightseeingPair(t *testing.T) {
	nums := [][]int{
		{8, 1, 5, 2, 6},
	}
	result := []int{
		11,
	}
	for i, r := range nums {
		rr := MaxScoreSightseeingPair(r)
		assert.Equal(t, rr, result[i])
	}
}
