package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxAreaOfLsLand(t *testing.T) {
	grids := [][][]int{
		{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}},
	}
	result := []int{
		4,
	}
	for i, r := range grids {
		rr := MaxAreaOfIsland(r)
		assert.Equal(t, rr, result[i])
	}

	grids = [][][]int{
		{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}},
	}
	for i, r := range grids {
		rr := MaxAreaOfIsland2(r)
		assert.Equal(t, rr, result[i])
	}
}
