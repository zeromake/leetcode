package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grids := [][][]byte{
		{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		nil,
	}
	result := []int{
		1,
		3,
		0,
	}
	for i, r := range grids {
		rr := NumIslands(r)
		assert.Equal(t, rr, result[i])
	}
}
