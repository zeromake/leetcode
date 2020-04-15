package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	ss := []string{
		"12",
		"226",
		"50",
		"101",
		"0111",
	}
	result := []int{
		2,
		3,
		0,
		1,
		0,
	}
	for i, r := range ss {
		rr := NumDecodings(r)
		assert.Equal(t, rr, result[i])
	}
}
