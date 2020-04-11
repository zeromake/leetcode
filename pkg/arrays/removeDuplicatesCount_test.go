package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesCount(t *testing.T) {
	nums := [][]int{
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	}
	result := [][]int{
		{1, 1, 2, 2, 3},
		{0, 0, 1, 1, 2, 3, 3},
	}
	for i, r := range nums {
		rr := RemoveDuplicatesCount(r)
		assert.Equal(t, r[:rr], result[i])
	}
}
