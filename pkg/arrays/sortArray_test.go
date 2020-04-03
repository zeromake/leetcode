package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := [][]int{
		{1, 2, 8, 5, 9},
		{9, 2, 8, 5, 1},
		{1, 2, 3, 3, 2, 1},
	}
	result := [][]int{
		{1, 2, 5, 8, 9},
		{1, 2, 5, 8, 9},
		{1, 1, 2, 2, 3, 3},
	}
	for i, r := range nums {
		rr := SortArray(r)
		assert.Equal(t, rr, result[i])
	}
}
