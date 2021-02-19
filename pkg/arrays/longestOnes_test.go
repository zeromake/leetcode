package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	nums := [][]int{
		{1,1,1,1,0,0,0,1,1,1,0},
		{1,1,1,0,0,0,1,1,1,1,0},
		{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1},
		nil,
	}
	k := []int{
		2,
		2,
		3,
		0,
	}
	result := []int{
		6,
		6,
		10,
		0,
	}
	for i, r := range nums {
		rr := LongestOnes(r, k[i])
		assert.Equal(t, rr, result[i])
	}
}
