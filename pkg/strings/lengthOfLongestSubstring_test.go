package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	res := []string{
		"abcabcbb",
		"",
	}
	result := []int{
		3,
		0,
	}
	for i, r := range res {
		assert.Equal(t, LengthOfLongestSubstring(r), result[i])
	}
}
