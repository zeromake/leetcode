package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T)  {
	coins := [][]int{
		{
			1,
			2,
			5,
		},
		{
			2,
		},
		{
		},
	}
	amounts := []int{
		11,
		3,
		5,
	}
	result := []int{
		3,
		-1,
		-1,
	}
	for i, c := range coins {
		r := CoinChange(c, amounts[i])
		assert.Equal(t, r, result[i])
	}
}

