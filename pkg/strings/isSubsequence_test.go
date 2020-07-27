package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	s := [][2]string{
		{
			"abc",
			"ahbgdc",
		},
		{
			"axc",
			"ahbgdc",
		},
		{
			"",
			"ahbgdc",
		},
		{
			"abcdefg",
			"ahbgdc",
		},
	}
	result := []bool{
		true,
		false,
		true,
		false,
	}
	for i, r := range s {
		rr := IsSubsequence(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
