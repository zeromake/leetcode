package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4},
	}
	result := [][]int{
		{24, 12, 8, 6},
	}
	for i, r := range nums {
		rr := ProductExceptSelf(r)
		assert.Equal(t, rr, result[i])
	}
}
