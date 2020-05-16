package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfitII(t *testing.T) {
	prices := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
	}
	result := []int{
		7,
		4,
	}
	for i, r := range prices {
		rr := MaxProfitII(r)
		assert.Equal(t, rr, result[i])
	}
}
