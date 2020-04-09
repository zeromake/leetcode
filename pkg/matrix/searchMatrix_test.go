package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][][]int{
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		{},
		{{}},
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
	}
	target := []int{
		3,
		1,
		1,
		0,
	}

	result := []bool{
		true,
		false,
		false,
		false,
	}
	for i, r := range matrix {
		rr := SearchMatrix(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}
