package strings

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAtoi(t *testing.T) {
	res := []string{
		" -42",
		"3.14159",
		"2000000000000000000",
		"+66",
		"-2000000000000000000",
		"7-",
		"7+",
		"7 ",
		"",
		"a",
	}
	result := []int{
		-42,
		3,
		math.MaxInt32,
		66,
		math.MinInt32,
		7,
		7,
		7,
		0,
		0,
	}
	for i, r := range res {
		assert.Equal(t, Atoi(r), result[i])
	}
}
