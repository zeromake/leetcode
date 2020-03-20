package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotatedSearch(t *testing.T) {
	nums := [][]int{
		{4, 5, 6, 7, 8, 1, 2, 3},
		{6, 7, 8, 1, 2, 3, 4, 5},
		{6, 7, 8, 1, 2, 3, 4, 5},
		{6, 7, 8, 1, 2, 3, 4, 5},
	}
	target := []int{
		5,
		6,
		5,
		9,
	}
	result := []int{
		1,
		0,
		7,
		-1,
	}
	for i, r := range nums {
		rr := RotatedSearch(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
