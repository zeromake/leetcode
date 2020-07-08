package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivingBoard(t *testing.T) {
	boards := [][3]int{
		{1, 2, 3},
		{1, 2, 0},
		{1, 2, 1},
		{1, 1, 3},
	}
	result := [][]int{
		{3, 4, 5, 6},
		nil,
		{1, 2},
		{3},
	}
	for i, r := range boards {
		rr := DivingBoard(r[0], r[1], r[2])
		assert.Equal(t, rr, result[i])
	}
}
