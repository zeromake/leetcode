package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfitIII(t *testing.T) {
	prices := [][]int{
		{
			3, 3, 5, 0, 3, 1, 4,
		},
	}
	result := []int{
		6,
	}
	for i, r := range prices {
		rr := MaxProfitIII(r)
		assert.Equal(t, rr, result[i])
	}
}
