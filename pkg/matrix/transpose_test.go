package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranspose(t *testing.T) {
	tickets := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		nil,
	}
	result := [][][]int{
		{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
		{
			{1, 4},
			{2, 5},
			{3, 6},
		},
		nil,
	}
	for i, r := range tickets {
		rr := Transpose(r)
		assert.Equal(t, rr, result[i])
	}
}
