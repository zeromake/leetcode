package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumbers(t *testing.T) {
	nums := [][]int{
		{4, 1, 4, 6},
		{1, 2, 10, 4, 1, 4, 3, 3},
	}
	result := [][]int{
		{1, 6},
		{10, 2},
	}
	for i, r := range nums {
		rr := SingleNumbers(r)
		assert.Equal(t, rr, result[i])
	}
}
