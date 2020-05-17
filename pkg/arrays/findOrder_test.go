package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOrder(t *testing.T) {
	infos := [][][]int{
		{
			{1, 0},
		},
		{
			{1, 0},
			{2, 0},
			{3, 1},
			{3, 2},
		},
		{
			{1, 0},
			{0, 1},
		},
		{
			{1, 0},
			{2, 1},
			{3, 2},
			{0, 3},
		},
	}
	nums := []int{
		2,
		4,
		2,
		4,
	}
	result := [][]int{
		{0, 1},
		{0, 2, 1, 3},
		nil,
		nil,
	}
	for i, r := range infos {
		rr := FindOrder(nums[i], r)
		assert.Equal(t, rr, result[i])
	}
}
