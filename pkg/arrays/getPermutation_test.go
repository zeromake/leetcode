package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	n := [][2]int{
		{3, 3},
		{4, 9},
	}
	result := []string{
		"213",
		"2314",
	}
	for i, r := range n {
		rr := GetPermutation(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
