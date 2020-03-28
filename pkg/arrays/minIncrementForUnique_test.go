package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinIncrementForUnique(t *testing.T) {
	nums := [][]int{
		{1, 2, 2},
		{2, 2, 2, 1},
		{1, 1, 1, 5},
	}
	result := []int{
		1,
		3,
		3,
	}
	for i, r := range nums {
		rr := MinIncrementForUnique(r)
		assert.Equal(t, rr, result[i])
		rr = MinIncrementForUnique2(r)
		assert.Equal(t, rr, result[i])
	}
}
