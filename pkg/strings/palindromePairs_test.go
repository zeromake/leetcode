package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromePairs(t *testing.T) {
	words := [][]string{
		{"abcd", "dcba", "lls", "s", "sssll"},
		{"bat", "tab", "cat"},
		{"bat", "tab", "cat", ""},
	}
	result := [][][]int{
		{
			{0, 1},
			{1, 0},
			{3, 2},
			{2, 4},
		},
		{
			{0, 1},
			{1, 0},
		},
		{
			{0, 1},
			{1, 0},
		},
	}
	for i, r := range words {
		rr := PalindromePairs(r)
		assert.Equal(t, rr, result[i])
	}
}
