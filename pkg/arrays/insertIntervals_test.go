package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertIntervals(t *testing.T) {
	intervals := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
	}

	news := [][]int{
		{11, 13},
		{11, 40},
	}

	result := [][][]int{
		{{1, 6}, {8, 10}, {11, 13}, {15, 18}},
		{{1, 10}, {11, 40}},
	}

	for i, r := range intervals {
		rr := InsertIntervals(r, news[i])
		assert.Equal(t, rr, result[i])
	}
}
