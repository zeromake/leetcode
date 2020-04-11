package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	n := [][2]int{
		{4, 2},
	}
	result := [][][]int{
		{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		},
	}
	for i, r := range n {
		rr := Combine(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = Combine2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
