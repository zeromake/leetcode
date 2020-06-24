package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	arr := [][][]int{
		{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		},
	}
	result := []int{
		11,
	}
	for i, r := range arr {
		rr := MinimumTotal(r)
		assert.Equal(t, rr, result[i])
	}
}
