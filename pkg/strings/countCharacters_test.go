package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	words := [][]string{
		{
			"cat",
			"bt",
			"help",
		},
	}
	chars := []string{
		"atcbu",
	}
	result := []int{
		5,
	}
	for i, r := range words {
		rr := CountCharacters(r, chars[i])
		assert.Equal(t, rr, result[i])
	}
}

