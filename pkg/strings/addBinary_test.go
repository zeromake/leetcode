package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	binarys := [][2]string{
		{"11", "1"},
		{"11", "11"},
		{"1", "11"},
	}
	result := []string{
		"100",
		"110",
		"100",
	}
	for i, r := range binarys {
		rr := AddBinary(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
