package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
	}
	result := []int{
		4,
		12,
	}
	for i, r := range nums {
		rr := Rob(r)
		assert.Equal(t, rr, result[i])
	}
}
