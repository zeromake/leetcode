package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := [][]int{
		{
			0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
		},
	}
	result := [][]int{
		{
			0, 1, 2, 3, 4,
		},
	}
	for i, r := range nums {
		rr := RemoveDuplicates(r)
		assert.Equal(t, rr, len(result[i]))
		assert.Equal(t, r[:rr], result[i])
	}
}
