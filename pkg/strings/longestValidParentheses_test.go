package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	ss := []string{
		"(()",
		"()()",
		"())()",
		"()(()",
	}
	result := []int{
		2,
		4,
		2,
		2,
	}
	for i, s := range ss {
		rr := LongestValidParentheses(s)
		assert.Equal(t, rr, result[i])
	}
}
