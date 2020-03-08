package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	res := [][][]int{
		{
			{
				2, 1, 1,
			},
			{
				1, 1, 0,
			},
			{
				0, 1, 1,
			},
		},
		{
			{
				2, 1, 1,
			},
			{
				0, 1, 1,
			},
			{
				1, 0, 1,
			},
		},
		{
			{
				1, 1, 1,
			},
			{
				1, 2, 1,
			},
			{
				1, 0, 1,
			},
		},
	}
	count := []int{
		4,
		-1,
		2,
	}
	for i, r := range res {
		rr := orangesRotting(r)
		assert.Equal(t, rr, count[i])
	}
	res = [][][]int{
		{
			{
				2, 1, 1,
			},
			{
				1, 1, 0,
			},
			{
				0, 1, 1,
			},
		},
		{
			{
				2, 1, 1,
			},
			{
				0, 1, 1,
			},
			{
				1, 0, 1,
			},
		},
		{
			{
				1, 1, 1,
			},
			{
				1, 2, 1,
			},
			{
				1, 0, 1,
			},
		},
	}
	for i, r := range res {
		rr := orangesRotting2(r)
		assert.Equal(t, rr, count[i])
	}
}
