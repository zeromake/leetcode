package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := [][]string{
		{
			"flower",
			"flow",
			"flight",
		},
		{
			"dog",
			"racecar",
			"car",
		},
		{
			"test1",
			"test",
		},
		{},
	}
	result := []string{
		"fl",
		"",
		"test",
		"",
	}
	for i, r := range strs {
		rr := LongestCommonPrefix(r)
		assert.Equal(t, rr, result[i])
	}
}
