package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	strings := [][2]string{
		{
			"98",
			"9",
		},
		{
			"9",
			"98",
		},
	}
	result := []string{
		"107",
		"107",
	}
	for i, r := range strings {
		rr := AddStrings(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
