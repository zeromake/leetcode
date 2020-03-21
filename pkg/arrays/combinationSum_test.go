package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	nums := [][]int{
		{2, 3, 6, 7},
		{2, 3, 5},
	}
	target := []int{
		7,
		8,
	}
	result := [][][]int{
		{
			{2, 2, 3,},
			{7},
		},
		{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		},
	}
	for i, r := range nums {
		rr := CombinationSum(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
func TestCombinationSum2(t *testing.T) {
	nums := [][]int{
		{2, 3, 6, 7},
		{2, 2, 3, 5},
	}
	target := []int{
		7,
		8,
	}
	result := [][][]int{
		{
			{7},
		},
		{
			{3, 5},
		},
	}
	for i, r := range nums {
		rr := CombinationSum2(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
