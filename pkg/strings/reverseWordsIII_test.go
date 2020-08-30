package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWordsIII(t *testing.T) {
	strs := []string{
		"Let's take LeetCode contest",
	}
	result := []string{
		"s'teL ekat edoCteeL tsetnoc",
	}
	for i, r := range strs {
		rr := ReverseWordsIII(r)
		assert.Equal(t, rr, result[i])
	}
}
