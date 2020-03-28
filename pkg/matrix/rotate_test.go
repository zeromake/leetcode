package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix(t *testing.T) {
	matrix := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{8, 9, 10, 11},
			{12, 13, 14, 15},
		},
	}
	result := [][][]int{
		{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
		{
			{12, 8, 5, 1},
			{13, 9, 6, 2},
			{14, 10, 7, 3},
			{15, 11, 8, 4},
		},
	}
	for i, r := range matrix {
		Rotate(r)
		assert.Equal(t, r, result[i])
	}
}
