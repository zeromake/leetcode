package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := [][3]string{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"ABCD",
			"ABCDE",
			"",
		},
	}
	for _, r := range s {
		rr := MinWindow(r[0], r[1])
		assert.Equal(t, rr, r[2])
	}
}
