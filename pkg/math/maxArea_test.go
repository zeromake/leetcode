package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T)  {
	heights := [][]int{
		{
			1, 8, 6, 2, 5, 4, 8, 3, 7,
		},
	}
	result := []int{
		49,
	}
	for i, h := range heights {
		area := MaxArea(h)
		assert.Equal(t, area, result[i])
	}
}

