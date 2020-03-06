package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type merge struct {
	A   []int
	M   int
	B   []int
	N   int
	Res []int
}

func TestMerge(t *testing.T) {
	res := []merge{
		merge{
			A: []int{
				2, 3, 0, 0, 0,
			},
			M: 2,
			B: []int{
				2, 5, 6,
			},
			N: 3,
			Res: []int{
				2, 2, 3, 5, 6,
			},
		},
		merge{
			A: []int{
				1, 0, 0, 0,
			},
			M: 1,
			B: []int{
				2, 5, 6,
			},
			N: 3,
			Res: []int{
				1, 2, 5, 6,
			},
		},
	}
	for _, r := range res {
		Merge(r.A, r.M, r.B, r.N)
		assert.Equal(t, r.Res, r.A, "error")
	}
}
