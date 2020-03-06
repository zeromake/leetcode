package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindContinuousSequence(t *testing.T)  {
	res := []int{
		9,
	}
	result := [][][]int{
		{
			{2, 3, 4},
			{4, 5},
		},
	}
	for i, r := range res {
		rr := FindContinuousSequence(r)
		assert.Equal(t, rr, result[i])
	}
}

