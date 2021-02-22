package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsToeplitzMatrix(t *testing.T) {
	recs := [][][]int{
		{
			{1, 2, 3, 4},
			{5, 1, 2, 3},
			{9, 5, 1, 2},
		},
		{
			{1, 2},
			{2, 2},
		},
		{
			{36, 59, 71, 15, 26, 82, 87},
			{56, 36, 59, 71, 15, 26, 82},
			{15, 0, 36, 59, 71, 15, 26},
		},
		{
			{1},
		},
		{
			nil,
		},
		nil,
	}
	result := []bool{
		true,
		false,
		false,
		true,
		false,
		false,
	}
	for i, r := range recs {
		rr := IsToeplitzMatrix(r)
		assert.Equal(t, rr, result[i])
	}
}
