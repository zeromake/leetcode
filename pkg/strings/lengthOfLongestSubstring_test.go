package strings

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	res := []string{
		"abcabcbb",
	}
	result := []int{
		3,
	}
	if LengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fail()
	}
}
