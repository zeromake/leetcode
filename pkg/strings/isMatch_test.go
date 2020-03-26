package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	strs := [][2]string{
		{
			"",
			"*",
		},
		{
			"aa",
			"a",
		},
		{
			"aa",
			"a*",
		},
		{
			"ab",
			".*",
		},
		{
			"aab",
			"c*a*b",
		},
	}
	result := []bool{
		false,
		false,
		true,
		true,
		true,
		false,
	}
	for i, r := range strs {
		rr := IsMatch(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = IsMatch2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
