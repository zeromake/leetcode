package arrays

import (
	"testing"
)

type resp struct {
	Nums   []int
	Target int
	Res    []int
}

func TestTwoSum(t *testing.T) {
	res := []resp{
		resp{
			[]int{
				2,
				7,
				11,
				15,
			},
			9,
			[]int{
				0,
				1,
			},
		},
		resp{
			[]int{},
			5,
			nil,
		},
	}
	for _, r := range res {
		rr := TwoSum(r.Nums, r.Target)
		if r.Res == nil && rr == nil {
			continue
		}
		if rr[0] != r.Res[0] || rr[1] != r.Res[1] {
			t.Fatal("error", rr, r.Res)
		}
	}
}
