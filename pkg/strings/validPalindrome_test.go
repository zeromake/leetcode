package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	strs := []string{
		"aba",
		"abaa",
		"abcc",
		"acbbca",
	}
	result := []bool{
		true,
		true,
		false,
		true,
	}
	for i, r := range strs {
		rr := ValidPalindrome(r)
		assert.Equal(t, rr, result[i])
	}
}
