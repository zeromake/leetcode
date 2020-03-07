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
	}
	result := []int{
		-42,
		3,
		math.MaxInt32,
	}
	for i, r := range res {
		assert.Equal(t, Atoi(r), result[i])
	}
}
