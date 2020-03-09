package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
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
	for i, r := range nums {
		rr := IntToRoman(r)
		assert.Equal(t, rr, romans[i])
		rr = IntToRoman2(r)
		assert.Equal(t, rr, romans[i])
	}
}
