package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	words := []string{
		"the sky is blue",
		"  hello word!  ",
	}
	result := []string{
		"blue is sky the",
		"word! hello",
	}
	for i, r := range words {
		rr := ReverseWords(r)
		assert.Equal(t, rr, result[i])
	}
}
