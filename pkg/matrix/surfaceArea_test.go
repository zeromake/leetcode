package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSurfaceArea(t *testing.T) {
	grid := [][][]int{
		{
			{2},
		},
		{
			{1, 2},
			{3, 4},
		},
		{
			{1, 0},
			{0, 2},
		},
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		},
	}
	result := []int{
		10,
		34,
		16,
		32,
		46,
	}
	for i, r := range grid {
		rr := SurfaceArea(r)
		assert.Equal(t, rr, result[i])
	}
}
