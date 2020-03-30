package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJump(t *testing.T) {
	nums := [][]int{
		{
			2,
			3,
			1,
			1,
			4,
		},
	}
	result := []int{
		2,
	}
	for i, r := range nums {
		rr := Jump(r)
		assert.Equal(t, rr, result[i])
	}
}
