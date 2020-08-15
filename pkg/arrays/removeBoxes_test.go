package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveBoxes(t *testing.T) {
	boxes := [][]int{
		{1, 3, 2, 2, 2, 3, 4, 3, 1},
		nil,
	}
	result := []int{
		23,
		0,
	}
	for i, r := range boxes {
		rr := RemoveBoxes(r)
		assert.Equal(t, rr, result[i])
	}
}
