package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	strs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		":A man nama;",
		"",
	}
	result := []bool{
		true,
		false,
		true,
		true,
	}
	for i, r := range strs {
		rr := IsPalindrome(r)
		assert.Equal(t, rr, result[i])
	}
}
