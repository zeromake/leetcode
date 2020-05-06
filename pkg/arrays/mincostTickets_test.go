package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCostTickets(t *testing.T) {
	nums := [][2][]int{
		{
			{1, 4, 6, 7, 8, 20},
			{2, 7, 15},
		},
		{
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			{2, 7, 15},
		},
	}
	result := []int{
		11,
		17,
	}
	for i, r := range nums {
		rr := MinCostTickets(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
