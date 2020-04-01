package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	intervals := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
		{},
	}

	result := [][][]int{
		{{1, 6}, {8, 10}, {15, 18}},
		{{1, 10}},
		nil,
	}

	for i, r := range intervals {
		rr := MergeIntervals(r)
		assert.Equal(t, rr, result[i])
		rr = MergeIntervals2(r)
		assert.Equal(t, rr, result[i])
	}
}
