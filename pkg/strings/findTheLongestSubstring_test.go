package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheLongestSubstring(t *testing.T) {
	strs := []string{
		"eleetminicoworoep",
		"leetcodeisgreat",
		"bcbcbc",
		"bcbcbcu",
	}
	result := []int{
		13,
		5,
		6,
		6,
	}
	for i, r := range strs {
		rr := FindTheLongestSubstring(r)
		assert.Equal(t, rr, result[i])
	}
}
