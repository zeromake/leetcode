package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastRemaining(t *testing.T) {
	nums := [][2]int{
		{5, 3},
		{7, 9},
		{10, 17},
	}
	result := []int{
		3,
		6,
		2,
	}
	for i, r := range nums {
		rr := LastRemaining(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
