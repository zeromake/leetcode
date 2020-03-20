package math

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	nums := [][2]int{
		{
			-1,
			-1,
		},
		{
			0,
			1,
		},
		{
			6,
			4,
		},
		{
			100,
			3,
		},
		{
			math.MinInt32 - 2,
			math.MinInt32>>1 - 1,
		},
		{
			math.MaxInt32 + 1,
			1,
		},
	}
	for _, r := range nums {
		rr := Divide(r[0], r[1])
		if rr != math.MaxInt32 {
			assert.Equal(t, rr, r[0]/r[1])
		}

		//rr = Divide2(r[0], r[1])
		//if rr != math.MaxInt32 {
		//	assert.Equal(t, rr, r[0] / r[1])
		//}
	}
}
