package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	nums := [][]int{
		{5, 2, 6, 1},
		nil,
	}
	result := [][]int{
		{2, 1, 1, 0},
		nil,
	}
	for i, r := range nums {
		rr := CountSmaller(r)
		assert.Equal(t, rr, result[i])
	}
}
