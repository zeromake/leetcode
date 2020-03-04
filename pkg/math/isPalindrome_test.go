package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	res := []int{
		1221,
		12321,
		-12321,
		0,
		10,
	}
	result := []bool{
		true,
		true,
		false,
		true,
		false,
	}
	for i, r := range res {
		rr := isPalindrome(r)
		assert.Equal(t, rr, result[i], r)
	}
}

