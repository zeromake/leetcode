package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSpiralMatrix(t *testing.T) {
	n := []int{
		2,
		3,
		4,
		5,
		6,
	}
	result := [][][]int{
		{
			{1, 2},
			{4, 3},
		},
		{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
		{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		},
		{
			{1, 2, 3, 4, 5},
			{16, 17, 18, 19, 6},
			{15, 24, 25, 20, 7},
			{14, 23, 22, 21, 8},
			{13, 12, 11, 10, 9},
		},
		{
			{1, 2, 3, 4, 5, 6},
			{20, 21, 22, 23, 24, 7},
			{19, 32, 33, 34, 25, 8},
			{18, 31, 36, 35, 26, 9},
			{17, 30, 29, 28, 27, 10},
			{16, 15, 14, 13, 12, 11},
		},
	}
	for i, r := range n {
		rr := GenerateSpiralMatrix(r)
		assert.Equal(t, rr, result[i])
	}
}
