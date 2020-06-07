package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLadders(t *testing.T) {
	words := [][2]string{
		{"hit", "cog"},
		{"hit", "cog"},
		{"", "cog"},
	}
	wordLists := [][]string{
		{"hot", "dot", "dog", "lot", "log", "cog"},
		{"hot", "dot", "dog", "lot", "log"},
		{"cog"},
	}
	result := [][][]string{
		{
			{"hit", "hot", "dot", "dog", "cog"},
			{"hit", "hot", "lot", "log", "cog"},
		},
		nil,
		nil,
	}
	for i, r := range words {
		rr := FindLadders(r[0], r[1], wordLists[i])
		assert.Equal(t, rr, result[i])
	}
}
