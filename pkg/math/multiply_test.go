package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	nums := [][2]string{
		{
			"10",
			"11",
		},
		{
			"11",
			"11",
		},
		{
			"3",
			"2",
		},
		{
			"123",
			"456",
		},
		{
			"0",
			"0",
		},
		{
			"0",
			"9133",
		},
	}
	result := []string{
		"110",
		"121",
		"6",
		"56088",
		"0",
		"0",
	}
	for i, r := range nums {
		rr := Multiply(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
