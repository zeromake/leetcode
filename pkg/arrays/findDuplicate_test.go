package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	nums := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
	}
	result := []int{
		2,
		3,
	}
	for i, r := range nums {
		rr := FindDuplicate(r)
		assert.Equal(t, rr, result[i])
	}
}
