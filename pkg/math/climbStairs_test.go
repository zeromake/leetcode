package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	n := []int{
		2,
		3,
		1,
	}
	result := []int{
		2,
		3,
		1,
	}
	for i, r := range n {
		rr := ClimbStairs(r)
		assert.Equal(t, rr, result[i])
	}
}
