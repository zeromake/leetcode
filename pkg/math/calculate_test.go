package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	ss := []string{
		"1+2-5",
		"4+6",
		"1+ 2 ",
		"2- 1",
		"3-(2 +1 )",
		"3+(1-2+1)",
	}
	result := []int{
		-2,
		10,
		3,
		1,
		0,
		3,
	}
	for i, r := range ss {
		rr := Calculate(r)
		assert.Equal(t, rr, result[i])
	}
}
