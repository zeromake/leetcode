package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsRectangleOverlap(t *testing.T) {
	recs := [][2][]int{
		{
			{
				0, 0, 2, 2,
			},
			{
				1, 1, 3, 3,
			},
		},
		{
			{
				0, 0, 1, 1,
			},
			{
				1, 0, 2, 1,
			},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range recs {
		rr := IsRectangleOverlap(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
