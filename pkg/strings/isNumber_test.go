package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNumber(t *testing.T) {
	s := []string{
		"0",
		"0.1",
		"abc",
		"1 a",
		"2e10",
		"-90e3",
		"1e",
		"e3",
		" 6e-1",
		"99e2.5",
		"53.5e93",
		" --6",
		"-+3",
		"95a54e53",
		"1e-",
		"1e.",
		"1e+1",
		"",
		".1",
		"2.",
		".",
		"..",
		".0e",
		"-1.",
		".-4",
		"-.",
	}
	result := []bool{
		true,
		true,
		false,
		false,
		true,
		true,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		false,
		false,
		true,
		false,
		true,
		true,
		false,
		false,
		false,
		true,
		false,
		false,
	}
	for i, r := range s {
		rr := IsNumber(r)
		assert.Equal(t, rr , result[i], r)
		rr = IsNumber2(r)
		assert.Equal(t, rr , result[i], r)
	}
}
