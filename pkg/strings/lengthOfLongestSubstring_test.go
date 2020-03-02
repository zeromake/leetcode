package strings

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	if LengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fail()
	}
}
