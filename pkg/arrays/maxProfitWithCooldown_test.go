package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfitWithCooldown(t *testing.T) {
	prices := [][]int{
		{1, 2, 3, 0, 2},
		nil,
	}
	result := []int{
		3,
		0,
	}
	for i, r := range prices {
		rr := MaxProfitWithCooldown(r)
		assert.Equal(t, rr, result[i])
	}
}
