package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovingCount(t *testing.T) {
	size := [][3]int{
		{3, 2, 17},
		{3, 2, 0},
	}
	result := []int{
		6,
		1,
	}
	for i, r := range size {
		rr := MovingCount(r[0], r[1], r[2])
		assert.Equal(t, rr, result[i])
	}
}

