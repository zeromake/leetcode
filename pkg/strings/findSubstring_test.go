package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	s := []string{
		"barfoothefoobarman",
		"wordgoodgoodgoodbestword",
		"",
	}
	words := [][]string{
		{
			"foo",
			"bar",
		},
		{
			"word",
			"good",
			"best",
			"good",
		},
		{},
	}
	result := [][]int{
		{0, 9},
		{8},
		nil,
	}
	for i, r := range s {
		rr := FindSubstring(r, words[i])
		assert.Equal(t, rr, result[i])

		rr = FindSubstring2(r, words[i])
		assert.Equal(t, rr, result[i])
	}
}

func BenchmarkFindSubstring(b *testing.B) {
	b.ReportAllocs()
	s := "wordgoodgoodgoodbestword"
	words := []string{
		"word",
		"good",
		"best",
		"good",
	}
	result := []int{8}
	for i := 0; i < b.N; i++ {
		r := FindSubstring(s, words)
		if len(r) != len(result) || r[0] != result[0] {
			b.Error(r, result)
		}
	}
}

func BenchmarkFindSubstring2(b *testing.B) {
	b.ReportAllocs()
	s := "wordgoodgoodgoodbestword"
	words := []string{
		"word",
		"good",
		"best",
		"good",
	}
	result := []int{8}
	for i := 0; i < b.N; i++ {
		r := FindSubstring2(s, words)
		if len(r) != len(result) || r[0] != result[0] {
			b.Error(r, result)
		}
	}
}
