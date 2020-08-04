package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanFinish(t *testing.T) {
	nums := []int{
		2,
		2,
	}
	prerequisites := [][][]int{
		{
			{1, 0},
		},
		{
			{1, 0},
			{0, 1},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range nums {
		rr := CanFinish(r, prerequisites[i])
		assert.Equal(t, rr, result[i])
	}
}
