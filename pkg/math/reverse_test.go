package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	result := [][2]int{
		{
			123,
			321,
		},
		{
			120,
			21,
		},
	}
	for _, r := range result {
		rr := Reverse(r[0])
		assert.Equal(t, rr, r[1])
	}
}
