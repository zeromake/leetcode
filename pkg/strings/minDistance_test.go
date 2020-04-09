package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDistance(t *testing.T) {
	words := [][2]string{
		{"horse", "ros"},
		{"intention", "execution"},
		{"", ""},
		{"", "cat"},
	}
	result := []int{
		3,
		5,
		0,
		3,
	}
	for i, r := range words {
		rr := MinDistance(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
