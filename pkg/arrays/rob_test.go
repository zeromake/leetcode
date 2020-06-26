package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{1},
		nil,
	}
	result := []int{
		4,
		12,
		1,
		0,
	}
	for i, r := range nums {
		rr := Rob(r)
		assert.Equal(t, rr, result[i])
	}
}
