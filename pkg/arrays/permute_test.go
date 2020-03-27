package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := [][]int{
		{2, 1, 3},
		{},
	}
	result := [][][]int{
		{
			{2, 1, 3},
			{2, 3, 1},
			{1, 2, 3},
			{1, 3, 2},
			{3, 2, 1},
			{3, 1, 2},
		},
		nil,
	}
	for i, r := range nums {
		rr := Permute(r)
		assert.Equal(t, rr, result[i])
	}
}


func TestPermuteUnique(t *testing.T) {
	nums := [][]int{
		{2, 1, 1, 3},
		{},
	}
	result := [][][]int{
		{
			{1, 1, 2, 3},
			{1, 1, 3, 2},
			{1, 2, 1, 3},
			{1, 2, 3, 1},
			{1, 3, 1, 2},
			{1, 3, 2, 1},
			{2, 1, 1, 3},
			{2, 1, 3, 1},
			{2, 3, 1, 1},
			{3, 1, 1, 2},
			{3, 1, 2, 1},
			{3, 2, 1, 1},
		},
		nil,
	}
	for i, r := range nums {
		rr := PermuteUnique(r)
		assert.Equal(t, rr, result[i])
	}
}
