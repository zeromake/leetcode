package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	nums := []uint32{
		43261596,
		4294967293,
	}
	result := []uint32{
		964176192,
		3221225471,
	}
	for i, r := range nums {
		rr := ReverseBits(r)
		assert.Equal(t, rr, result[i])
	}
}
