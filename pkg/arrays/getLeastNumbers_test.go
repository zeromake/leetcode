package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1},
		{1, 2},
	}
	k := []int{
		2,
		2,
		2,
		0,
	}
	result := [][]int{
		{2, 1},
		{2, 1},
		{1},
		nil,
	}
	for i, r := range nums {
		rr := GetLeastNumbers(r, k[i])
		assert.Equal(t, rr, result[i])
	}
}
