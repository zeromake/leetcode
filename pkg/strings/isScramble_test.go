package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsScramble(t *testing.T) {
	strs := [][2]string{
		{"rgeat", "great"},
		{"rgtae", "great"},
		{"abcde", "caebd"},
		{"abcde", "caebdf"},
	}
	result := []bool{
		true,
		true,
		false,
		false,
	}
	for i, r := range strs {
		rr := IsScramble(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = IsScramble2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
