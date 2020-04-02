package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	n := []int{
		4,
	}
	result := []int{
		2,
	}
	for i, r := range n {
		rr := TotalNQueens(r)
		assert.Equal(t, rr, result[i])
	}
}

