package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	ss := [][2]string{
		{
			"ABCABC",
			"ABC",
		},
		{
			"ABABAB",
			"ABAB",
		},
		{
			"LEFT",
			"CODE",
		},
	}
	result := []string{
		"ABC",
		"AB",
		"",
	}
	for i, r := range ss {
		rr := GcdOfStrings(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = GcdOfStrings2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
