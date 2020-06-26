package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	boards := [][][]byte{
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		nil,
	}
	result := [][][]byte{
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		nil,
	}
	for i, r := range boards {
		Solve(r)
		assert.Equal(t, r, result[i])
	}
}
