package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	x := []int{
		4,
		8,
		1,
		0,
	}
	result := []int{
		2,
		2,
		1,
		0,
	}
	for i, r := range x {
		rr := MySqrt(r)
		assert.Equal(t, rr, result[i])
		rr = MySqrt2(r)
		assert.Equal(t, rr, result[i])
	}
}
