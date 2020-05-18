package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := [][]int{
		{2, 3, -2, 4},
		{-2, 0, -1},
	}
	result := []int{
		6,
		0,
	}
	for i, r := range nums {
		rr := MaxProduct(r)
		assert.Equal(t, rr, result[i])
	}
}
