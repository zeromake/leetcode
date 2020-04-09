package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	size := [][2]int{
		{3, 2},
		{7, 3},
		{23, 12},
	}
	result := []int{
		3,
		28,
		193536720,
	}
	for i, r := range size {
		rr := UniquePaths(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = UniquePaths2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
