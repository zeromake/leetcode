package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := [][2][]int{
		{
			{1, 2, 3, 0, 0, 0},
			{2, 5, 6},
		},
		{
			{14, 15, 16, 0, 0, 0},
			{2, 5, 6},
		},
	}
	n := [][2]int{
		{3, 3},
		{3, 3},
	}
	result := [][]int{
		{1, 2, 2, 3, 5, 6},
		{2, 5, 6, 14, 15, 16},
	}
	for i, r := range nums {
		nm := n[i]
		MergeSort(r[0], nm[0], r[1], nm[1])
		assert.Equal(t, r[0], result[i])
	}
}
