package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type median struct {
	A   []int
	B   []int
	Res float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	res := []median{
		median{
			[]int{
				1, 3,
			},
			[]int{
				2,
			},
			2.0,
		},
		median{
			[]int{
				1, 3,
			},
			[]int{
				2, 4,
			},
			2.5,
		},
		median{
			[]int{
				1, 2,
			},
			[]int{
				3, 4,
			},
			2.5,
		},
		median{
			[]int{
				2,
			},
			[]int{},
			2.0,
		},
		median{
			[]int{
				1, 1, 3, 3,
			},
			[]int{
				1, 1, 3, 3,
			},
			2.0,
		},
		median{
			[]int{
				1,
			},
			[]int{
				2, 3, 4, 5,
			},
			3.0,
		},
	}
	for _, r := range res {
		rr := FindMedianSortedArrays(r.A, r.B)
		assert.Equal(t, rr, r.Res)

		rr = FindMedianSortedArrays2(r.A, r.B)
		assert.Equal(t, rr, r.Res)

		rr = FindMedianSortedArrays3(r.A, r.B)
		assert.Equal(t, rr, r.Res)
	}
}
