package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	res := [][2]int{
		{
			7,
			4,
		},
		{
			10,
			3,
		},
	}
	result := [][]int{
		{
			1,2,3,1,
		},
		{
			5,2,3,
		},
	}
	for i, r := range res {
		rr := DistributeCandies(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
