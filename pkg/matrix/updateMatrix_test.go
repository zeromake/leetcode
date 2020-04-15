package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	matrix := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
	}
	result := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		},
	}
	for i, r := range matrix {
		rr := UpdateMatrix(r)
		assert.Equal(t, rr, result[i])
	}
}
