package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	nums := [][]int{
		{1, 5, 2},
		{1, 5, 233, 7},
	}
	result := []bool{
		false,
		true,
	}
	for i, r := range nums {
		rr := PredictTheWinner(r)
		assert.Equal(t, rr, result[i])
	}
}
