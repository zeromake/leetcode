package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	nums := []float64{
		2.0,
		2.1,
		2.0,
		5.1,
	}
	pow := []int{
		10,
		3,
		-2,
		0,
	}
	result := []float64{
		1024,
		9.261000000000001,
		0.25,
		1,
	}
	for i, r := range nums {
		rr := MyPow(r, pow[i])
		assert.Equal(t, rr, result[i])
	}
}
