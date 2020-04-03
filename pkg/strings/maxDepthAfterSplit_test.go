package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepthAfterSplit(t *testing.T) {
	seqs := []string{
		"(()())",
		"()(())()",
	}
	result := [][]int{
		{0, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
	}
	for i, r := range seqs {
		rr := MaxDepthAfterSplit(r)
		assert.Equal(t, rr, result[i])
	}
}
