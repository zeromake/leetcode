package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		},
		{
			{6, 9, 7},
		},
		{
			{7}, {9}, {6},
		},
		nil,
	}
	result := [][]int{
		{1, 2, 3, 6, 9, 8, 7, 4, 5},
		{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		{6, 9, 7},
		{7, 9, 6},
		nil,
	}
	for i, r := range matrix {
		rr := SpiralOrder(r)
		assert.Equal(t, rr, result[i])
	}
}
