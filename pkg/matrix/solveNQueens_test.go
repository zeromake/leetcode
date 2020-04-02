package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	n := []int{
		4,
	}
	result := [][][]string{
		{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
	}
	for i, r := range n {
		rr := SolveNQueens(r)
		assert.Equal(t, rr, result[i])
		rr = SolveNQueens2(r)
		assert.Equal(t, rr, result[i])
	}
}
