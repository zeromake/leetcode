package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew21Game(t *testing.T) {
	nums := [][3]int{
		{10, 1, 10},
		{6, 1, 10},
		{21, 17, 10},
		{10, 0, 10},
	}
	result := []float64{
		1.0,
		0.6,
		0.7327777870686083,
		1.0,
	}
	for i, r := range nums {
		rr := New21Game(r[0], r[1], r[2])
		assert.Equal(t, rr, result[i])
	}
}
