package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	res := []resp{
		{
			[]int{
				2,
				7,
				11,
				15,
			},
			9,
			[]int{
				1,
				2,
			},
		},
		{
			[]int{},
			5,
			[]int{
				-1,
				-1,
			},
		},
		{
			[]int{
				2,
				7,
				11,
				15,
			},
			18,
			[]int{
				2,
				3,
			},
		},
	}
	for _, r := range res {
		rr := TwoSumII(r.Nums, r.Target)
		assert.Equal(t, rr, r.Res)
	}
}
