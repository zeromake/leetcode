package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHappy(t *testing.T) {
	nums := []int{
		19,
		116,
	}
	result := []bool{true, false}
	for i, r := range nums {
		rr := IsHappy(r)
		assert.Equal(t, rr, result[i])
	}
}
