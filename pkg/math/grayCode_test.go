package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrayCode(t *testing.T) {
	n := []int{
		0,
		2,
		3,
	}
	result := [][]int{
		{0},
		{0, 1, 3, 2},
		{0, 1, 3, 2, 6, 7, 5, 4},
	}
	for i, r := range n {
		rr := GrayCode(r)
		assert.Equal(t, rr, result[i])
	}
}
