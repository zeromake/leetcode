package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := [][]int{
		{
			7,
			1,
			5,
			3,
			6,
			4,
		},
		{
			2,
			1,
			4,
		},
	}
	result := []int{
		5,
		3,
	}
	for i, r := range prices {
		rr := MaxProfit(r)
		assert.Equal(t, rr, result[i])
	}
}
