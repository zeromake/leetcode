package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 1, 1, 2},
		nil,
	}
	result := []int{
		4,
		3,
		0,
	}
	for i, r := range nums {
		rr := LongestConsecutive2(r)
		assert.Equal(t, rr, result[i])

		rr = LongestConsecutive(r)
		assert.Equal(t, rr, result[i])
	}
}
