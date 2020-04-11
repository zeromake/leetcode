package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotatedSearchDuplicate(t *testing.T) {
	nums := [][]int{
		{3, 1},
		{1, 3, 1, 1, 1},
		{2, 5, 6, 0, 0, 1, 2},
		{2, 5, 6, 0, 0, 1, 2},
		{4, 5, 6, 7, 8, 1, 2, 3},
		{6, 7, 8, 1, 2, 3, 4, 5},
		{6, 7, 8, 1, 2, 3, 4, 5},
		{6, 7, 8, 1, 2, 3, 4, 5},
	}
	target := []int{
		1,
		3,
		0,
		3,
		5,
		6,
		5,
		9,
	}
	result := []bool{
		true,
		true,
		true,
		false,
		true,
		true,
		true,
		false,
	}
	for i, r := range nums {
		rr := RotatedSearchDuplicate(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
