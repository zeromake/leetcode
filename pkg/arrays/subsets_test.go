package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
	}
	result := [][][]int{
		{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 3},
			{2},
			{2, 3},
			{3},
		},
	}
	for i, r := range nums {
		rr := Subsets(r)
		assert.Equal(t, rr, result[i])
	}
}
