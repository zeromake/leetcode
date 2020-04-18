package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	matrix := [][][]byte{
		{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		nil,
	}
	result := []int{
		6,
		0,
	}
	for i, r := range matrix {
		rr := MaximalRectangle(r)
		assert.Equal(t, rr, result[i])
	}
}
