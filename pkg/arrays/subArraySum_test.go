package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubArraySum(t *testing.T) {
	nums := [][]int{
		{1, 1, 1},
		{2, 3, 4, 5, 6},
	}
	k := []int{
		2,
		7,
	}
	result := []int{
		2,
		1,
	}
	for i, r := range nums {
		rr := SubArraySum(r, k[i])
		assert.Equal(t, rr, result[i])
	}
}
