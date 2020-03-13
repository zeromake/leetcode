package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := [][]int{
		{
			3, 2, 3,
		},
		{
			2, 2, 1, 1, 2, 2,
		},
	}
	result := []int{
		3,
		2,
	}
	for i, r := range nums {
		rr := MajorityElement(r)
		assert.Equal(t, rr, result[i])

		rr = MajorityElement2(r)
		assert.Equal(t, rr, result[i])
	}
}
