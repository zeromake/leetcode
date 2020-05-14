package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := [][]int{
		{2, 2, 1},
		{1, 1, 3, 3, 2},
	}
	result := []int{
		1,
		2,
	}
	for i, r := range nums {
		rr := SingleNumber(r)
		assert.Equal(t, rr, result[i])
	}
}
