package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	arr := [][]int{
		{2, 3, 1, 2, 4, 3},
		{1},
		nil,
	}
	s := []int{
		7,
		2,
		2,
	}
	result := []int{
		2,
		0,
		0,
	}
	for i, r := range arr {
		rr := MinSubArrayLen(s[i], r)
		assert.Equal(t, rr, result[i])
	}
}
