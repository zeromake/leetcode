package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := [][]int{
		{
			1, 2, 3, 4, 5, 6,
		},
		{
			5, 4, 3, 2, 1,
		},
		{
			1, 2, 3, 8, 5, 7, 6, 4,
		},
		{
			1,
		},
	}
	result := [][]int{
		{
			1, 2, 3, 4, 6, 5,
		},
		{
			1, 2, 3, 4, 5,
		},
		{
			1, 2, 3, 8, 6, 4, 5, 7,
		},
		{
			1,
		},
	}
	for i, r := range nums {
		NextPermutation(r)
		assert.Equal(t, r, result[i])
	}
}
