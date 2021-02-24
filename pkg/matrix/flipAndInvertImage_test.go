package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	tickets := [][][]int{
		{
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
		{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 1, 0},
		},
	}
	result := [][][]int{
		{
			{1, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 1, 0},
		},
	}
	for i, r := range tickets {
		rr := FlipAndInvertImage(r)
		assert.Equal(t, rr, result[i])
	}
}
