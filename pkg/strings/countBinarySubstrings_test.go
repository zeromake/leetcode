package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinarySubstrings(t *testing.T) {
	s := []string{
		"00110011",
		"10101",
	}
	result := []int{
		6,
		4,
	}
	for i, r := range s {
		rr := CountBinarySubstrings(r)
		assert.Equal(t, rr, result[i])
	}
}
