package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := [][]string{
		{
			"This",
			"is",
			"an",
			"example",
			"of",
			"text",
			"justification.",
		},
		{
			"What",
			"must",
			"be",
			"acknowledgment",
			"shall",
			"be",
		},
		{
			"Science",
			"is",
			"what",
			"we",
			"understand",
			"well",
			"enough",
			"to",
			"explain",
			"to",
			"a",
			"computer.",
			"Art",
			"is",
			"everything",
			"else",
			"we",
			"do",
		},
	}
	widths := []int{
		16, 16, 20,
	}
	result := [][]string{
		{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
		{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
		{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	}
	for i, r := range words {
		rr := FullJustify(r, widths[i])
		assert.Equal(t, rr, result[i])
	}
}
