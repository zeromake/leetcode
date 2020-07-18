package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	graph := [][][]int{
		{
			{1, 3},
			{0, 2},
			{1, 3},
			{0, 2},
		},
		{
			{1, 2, 3},
			{0, 2},
			{0, 1, 3},
			{0, 2},
		},
	}
	result := []bool{
		true,
		false,
	}
	for i, r := range graph {
		rr := IsBipartite(r)
		assert.Equal(t, rr, result[i])
	}
}
