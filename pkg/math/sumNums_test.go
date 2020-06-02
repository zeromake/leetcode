package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNums(t *testing.T) {
	nums := []int{
		3,
		4,
		5,
	}
	result := []int{
		6,
		10,
		15,
	}
	for i, r := range nums {
		rr := SumNums(r)
		assert.Equal(t, rr, result[i])
	}
}
