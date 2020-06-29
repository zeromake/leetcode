package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	arr := [][]int{
		{3, 2, 1, 5, 6, 4},
		{3, 2, 3, 1, 2, 4, 5, 5, 6},
	}
	k := []int{
		2,
		4,
	}
	result := []int{
		5,
		4,
	}
	for i, r := range arr {
		rr := FindKthLargest(r, k[i])
		assert.Equal(t, rr, result[i])
	}
}
