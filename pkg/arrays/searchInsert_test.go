package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := [][]int{
		{1, 3},
		{1, 3, 5, 7},
		{1, 3, 5, 7},
		{1, 3, 5, 7},
		{1, 3, 5, 7},
		{1, 3, 5, 7},
	}
	target := []int{
		2,
		2,
		5,
		4,
		8,
		0,
	}
	result := []int{
		1,
		1,
		2,
		2,
		4,
		0,
	}
	for i, r := range nums {
		rr := SearchInsert(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
