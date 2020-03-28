package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMassage(t *testing.T) {
	nums := [][]int{
		{
			1, 2, 3, 1,
		},
		{
			2, 7, 9, 3, 1,
		},
		{
			2, 1, 4, 5, 3, 1, 1, 3,
		},
		{},
		{
			5,
		},
	}
	result := []int{
		4,
		12,
		12,
		0,
		5,
	}
	for i, r := range nums {
		rr := Massage(r)
		assert.Equal(t, rr, result[i])
		rr = Massage2(r)
		assert.Equal(t, rr, result[i])
		rr = Massage3(r)
		assert.Equal(t, rr, result[i])
	}
}
