package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4, 1},
		{1},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range nums {
		rr := ContainsDuplicate(r)
		assert.Equal(t, rr, result[i])
	}
}

