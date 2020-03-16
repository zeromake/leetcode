package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompressString(t *testing.T) {
	strs := []string{
		"aabcccccaa",
		"",
		"bb",
	}
	result := []string{
		"a2b1c5a2",
		"",
		"bb",
	}
	for i, r := range strs {
		rr := CompressString(r)
		assert.Equal(t, rr, result[i])
	}
}
