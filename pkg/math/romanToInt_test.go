package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	nums := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	romans := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
	}
	for i, r := range romans {
		rr := RomanToInt(r)
		assert.Equal(t, rr, nums[i])
	}
}
