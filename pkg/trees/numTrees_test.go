package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTrees(t *testing.T) {
	n := []int{
		3,
	}
	result := []int{
		5,
	}
	for i, r := range n {
		rr := NumTrees(r)
		assert.Equal(t, rr, result[i])
	}
}
