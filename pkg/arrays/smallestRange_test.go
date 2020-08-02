package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRange(t *testing.T) {
	arr := [][][]int{
		{
			{4, 10, 15, 24, 26},
			{0, 9, 12, 20},
			{5, 18, 22, 30},
		},
	}
	result := [][]int{
		{
			20,
			24,
		},
	}
	for i, r := range arr {
		rr := SmallestRange(r)
		assert.Equal(t, rr, result[i])
	}
}
