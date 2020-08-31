package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	rooms := [][][]int{
		{
			{1},
			{2},
			{3},
			nil,
		},
		{
			{1, 3},
			{3, 0, 1},
			{2},
			{0},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range rooms {
		rr := CanVisitAllRooms(r)
		assert.Equal(t, rr, result[i])
	}
}
