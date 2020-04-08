package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	matrix := [][][]int{
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		},
	}
	result := [][][]int{
		{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		},
		{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		},
	}
	for i, r := range matrix {
		SetZeroes(r)
		assert.Equal(t, r, result[i])
	}
}
