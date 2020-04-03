package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	grids := [][][]int{
		{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
		{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		nil,
	}
	result := []int{
		-1,
		2,
		4,
		-1,
	}
	for i, r := range grids {
		rr := MaxDistance(r)
		assert.Equal(t, rr, result[i])
	}
}
