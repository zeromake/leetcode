package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTriangle(t *testing.T) {
	rows := []int{
		0,
		1,
		3,
	}
	result := [][][]int{
		nil,
		{
			{1},
		},
		{
			{1},
			{1, 1},
			{1, 2, 1},
		},
	}
	for i, r := range rows {
		rr := GenerateTriangle(r)
		assert.Equal(t, rr, result[i])
	}
}
