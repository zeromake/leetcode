package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindInMountainArray(t *testing.T) {
	nums := []*MountainArray{
		{
			Array: []int{1, 2, 3, 4, 5, 3, 1},
		},
		{
			Array: []int{1, 3, 4, 5, 3, 2, 1},
		},
		{
			Array: []int{0, 1, 2, 4, 2, 1},
		},
	}
	target := []int{
		3,
		2,
		3,
	}
	result := []int{
		2,
		5,
		-1,
	}
	for i, r := range nums {
		rr := FindInMountainArray(target[i], r)
		assert.Equal(t, rr, result[i])
	}
}
