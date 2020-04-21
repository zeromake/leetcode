package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfSubarrays(t *testing.T) {
	nums := [][]int{
		{1, 1, 2, 1, 1},
		{2, 4, 6},
		{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
	}
	k := []int{
		3,
		1,
		2,
	}
	result := []int{
		2,
		0,
		16,
	}
	for i, r := range nums {
		kk := k[i]
		rr := NumberOfSubarrays(r, kk)
		assert.Equal(t, rr, result[i])
	}
}
