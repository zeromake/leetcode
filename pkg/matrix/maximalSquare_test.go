package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrixs := [][][]byte{
		{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		nil,
	}
	result := []int{
		4,
		0,
	}
	for i, r := range matrixs {
		rr := MaximalSquare(r)
		assert.Equal(t, rr, result[i])
	}
}
