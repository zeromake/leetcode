package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := [][]int{
		{
			10,
			2,
			5,
			3,
			7,
			1,
			101,
			18,
		},
		{
			1,
		},
	}
	result := []int{
		4,
		1,
	}
	for i, r := range nums {
		rr := LengthOfLIS(r)
		assert.Equal(t, rr, result[i])
		rr = LengthOfLIS2(r)
		assert.Equal(t, rr, result[i])
	}
}

