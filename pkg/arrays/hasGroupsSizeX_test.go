package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasGroupsSizeX(t *testing.T) {
	deck := [][]int{
		{1, 2, 3, 4, 4, 3, 2, 1},
		{1, 1, 1, 2, 2, 2, 3, 3},
		{1},
		{1, 1},
		{1, 1, 2, 2, 2, 2},
		{1,1,1,1,2,2,2,2,2,2},
		{1,1,1,1,1,1,2,2,2,2,2,2,2,2,2,3,3,3,3,3,3,3,3},
		{},
	}
	result := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
		false,
		false,
	}
	for i, r := range deck {
		rr := HasGroupsSizeX(r)
		assert.Equal(t, rr, result[i], i)
	}
}

