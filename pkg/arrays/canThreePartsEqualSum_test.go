package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	arrs := [][]int{
		{
			18,
			12,
			-18,
			18,
			-19,
			-1,
			10,
			10,
		},
	}
	result := []bool{
		true,
	}
	for i, r := range arrs {
		rr := CanThreePartsEqualSum(r)
		assert.Equal(t, rr, result[i])
	}
}


