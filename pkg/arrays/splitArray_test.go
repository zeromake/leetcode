package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitArray(t *testing.T) {
	nums := [][]int{
		{7, 2, 5, 10, 8},
	}
	m := []int{
		2,
	}
	result := []int{
		18,
	}
	for i, r := range nums {
		rr := SplitArray(r, m[i])
		assert.Equal(t, rr, result[i])
	}
}
