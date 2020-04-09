package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := [][]int{
		{2, 0, 2, 1, 1, 0},
		{1},
	}
	result := [][]int{
		{0, 0, 1, 1, 2, 2},
		{1},
	}
	for i, r := range nums {
		SortColors(r)
		assert.Equal(t, r, result[i])
	}
}
