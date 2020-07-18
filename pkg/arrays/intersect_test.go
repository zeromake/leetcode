package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	nums := [][2][]int{
		{
			{1, 2, 2, 1},
			{2, 2},
		},
		{
			{4, 9, 5},
			{9, 4, 9, 8, 4},
		},
	}
	result := [][][]int{
		{
			{2, 2},
		},
		{
			{4, 9},
			{9, 4},
		},
	}
	for i, r := range nums {
		rr := Intersect(r[0], r[1])
		flag := false
		for _, res := range result[i] {
			flag = flag || assert.ObjectsAreEqual(rr, res)
		}
		assert.Equal(t, flag, true)
	}
}
