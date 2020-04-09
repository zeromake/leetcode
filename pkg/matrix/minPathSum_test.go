package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
	}
	result := []int{
		7,
	}
	for i, r := range grid {
		rr := MinPathSum(r)
		assert.Equal(t, rr, result[i])
	}
}
