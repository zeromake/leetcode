package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxCoins(t *testing.T) {
	nums := [][]int{
		{3, 1, 5, 8},
	}
	result := []int{
		167,
	}
	for i, r := range nums {
		rr := MaxCoins(r)
		assert.Equal(t, rr, result[i])
	}
}
