package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T)  {
	strings := []string{
		"23",
		"023",
	}
	result := [][]string{
		{
			"ad",
			"bd",
			"cd",
			"ae",
			"be",
			"ce",
			"af",
			"bf",
			"cf",
		},
		{
			"ad",
			"bd",
			"cd",
			"ae",
			"be",
			"ce",
			"af",
			"bf",
			"cf",
		},
	}
	for i, r := range strings {
		rr := LetterCombinations(r)
		assert.Equal(t, rr, result[i])
	}
}

