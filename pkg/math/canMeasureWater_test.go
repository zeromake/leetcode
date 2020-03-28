package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanMeasureWater(t *testing.T) {
	waters := [][3]int{
		{3, 5, 4},
		{11, 13, 6},
		{22003, 31237, 1},
		{0, 2, 2},
		{0, 2, 0},
		{1, 2, 4},
	}
	result := []bool{
		true,
		true,
		true,
		true,
		true,
		false,
	}
	for i, r := range waters {
		rr := CanMeasureWater(r[0], r[1], r[2])
		assert.Equal(t, rr, result[i], i)
	}
}
