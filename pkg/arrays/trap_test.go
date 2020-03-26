package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap(t *testing.T) {
	height := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}
	result := []int{
		6,
	}
	for i, r := range height {
		rr := Trap(r)
		assert.Equal(t, rr, result[i])
	}
}
