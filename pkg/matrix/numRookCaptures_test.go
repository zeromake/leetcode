package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumRookCaptures(t *testing.T) {
	board := [][][]byte{
		{
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', 'R', '.', '.', '.', 'p'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
		},
		{
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
			{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
			{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
			{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
			{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
		},
		{
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', 'B', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'},
		},
	}
	result := []int{
		3,
		0,
		3,
	}
	for i, r := range board {
		rr := NumRookCaptures(r)
		assert.Equal(t, rr, result[i])
	}
}
