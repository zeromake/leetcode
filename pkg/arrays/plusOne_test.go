package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
		{9},
		{1, 9, 9},
	}
	result := [][]int{
		{1, 2, 4},
		{1, 0},
		{2, 0, 0},
	}
	for i, r := range nums {
		rr := PlusOne(r)
		assert.Equal(t, rr, result[i])
	}
}

