package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumLengthEncoding(t *testing.T) {
	words := [][]string{
		{"time", "me", "bell"},
		{"atime", "me", "bell"},
		{"atime", "me", "bell", "bbell"},
	}
	result := []int{
		10,
		11,
		12,
	}
	for i, r := range words {
		rr := MinimumLengthEncoding(r)
		assert.Equal(t, rr, result[i])
	}
}
