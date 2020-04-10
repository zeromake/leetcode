package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][][]byte{
		{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		{
			{'a', 'a'},
		},
		{
			{'b'},
			{'a'},
			{'b'},
		},
		{
			{'a', 'b'},
			{'c', 'd'},
		},
		{
			{'a'},
		},
	}
	word := []string{
		"ABCCED",
		"SEE",
		"ABCB",
		"aaa",
		"bbabab",
		"abcd",
		"a",
	}
	result := []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		true,
	}
	for i, r := range board {
		rr := Exist(r, word[i])
		assert.Equal(t, rr, result[i])
		rr = Exist2(r, word[i])
		assert.Equal(t, rr, result[i])
	}
}


