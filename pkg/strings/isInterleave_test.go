package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	strs := [][3]string{
		{
			"aabcc",
			"dbbca",
			"aadbbcbcac",
		},
		{
			"aabcc",
			"dbbca",
			"aadbbbaccc",
		},
		{
			"aabcc",
			"dbbca",
			"aadbbbacccz",
		},
	}
	result := []bool{
		true,
		false,
		false,
	}
	for i, r := range strs {
		rr := IsInterleave(r[0], r[1], r[2])
		assert.Equal(t, rr, result[i])
	}
}
