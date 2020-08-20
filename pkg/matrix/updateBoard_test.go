package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	boards := [][][]byte{
		{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		},
		{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	}
	clicks := [][]int{
		{3, 0},
		{1, 2},
	}
	result := [][][]byte{
		{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
		{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'X', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	}
	for i, r := range boards {
		rr := UpdateBoard(r, clicks[i])
		assert.Equal(t, rr, result[i])
	}
}
