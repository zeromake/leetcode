package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	nums := [][]int{
		{
			3, 2, 2, 3,
		},
	}
	result := [][]int{
		{
			2, 2,
		},
	}
	for i, r := range nums {
		rr := RemoveElement(r, 3)
		assert.Equal(t, rr, len(result[i]))
		assert.Equal(t, r[:rr], result[i])
	}
}
