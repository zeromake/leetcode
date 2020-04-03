package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	board := [][][]int{
		{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
		{},
	}
	result := [][][]int{
		{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		{},
	}
	for i, r := range board {
		GameOfLife(r)
		assert.Equal(t, r, result[i])
	}
}
