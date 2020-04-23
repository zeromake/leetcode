package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaysToChange(t *testing.T) {
	n := []int{
		5,
		10,
	}
	result := []int{
		2,
		4,
	}
	for i, r := range n {
		rr := WaysToChange(r)
		assert.Equal(t, rr, result[i])
	}
}
