package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 5},
		{2, 3, 4},
		{1, 2, 3, 4},
	}
	result := []int{
		4,
		1,
		5,
	}
	for i, r := range nums {
		rr := FirstMissingPositive(r)
		assert.Equal(t, rr, result[i])
	}
}

