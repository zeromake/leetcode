package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	nums := [][]int{
		{
			2,
			3,
			1,
			1,
			4,
		},
		{
			3,
			2,
			1,
			0,
			4,
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range nums {
		rr := CanJump(r)
		assert.Equal(t, rr, result[i])
	}
}
