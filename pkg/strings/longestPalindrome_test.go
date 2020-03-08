package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	res := [][2]string{
		{
			"babad",
			"aba",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"",
			"",
		},
	}
	for _, r := range res {
		rr := LongestPalindrome(r[0])
		assert.Equal(t, rr, r[1])
	}
}
