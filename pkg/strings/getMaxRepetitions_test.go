package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxRepetitions(t *testing.T) {
	strs := [][2]string{
		{
			"acb",
			"ab",
		},
		{
			"abaacdbac",
			"adcbd",
		},
		{
			"ab",
			"ab",
		},
		{"abc", "ab"},
	}
	n := [][2]int{
		{
			4,
			2,
		},
		{
			100,
			4,
		},
		{
			2,
			2,
		},
		{
			0,
			2,
		},
	}
	result := []int{
		2,
		12,
		1,
		0,
	}
	for i, r := range strs {
		rr := GetMaxRepetitions(r[0], n[i][0], r[1], n[i][1])
		assert.Equal(t, rr, result[i])
	}
}
