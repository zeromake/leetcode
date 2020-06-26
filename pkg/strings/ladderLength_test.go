package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadderLength(t *testing.T) {
	words := [][2]string{
		{"hit", "cog"},
		{"hit", "cog"},
		{"hot", "dog"},
	}
	wordList := [][]string{
		{"hot", "dot", "dog", "lot", "log", "cog"},
		{"hot", "dot", "dog", "lot", "log"},
		{"hot", "dog"},
	}
	result := []int{
		5,
		0,
		0,
	}
	for i, r := range words {
		rr := LadderLength(r[0], r[1], wordList[i])
		assert.Equal(t, rr, result[i])
	}
}
