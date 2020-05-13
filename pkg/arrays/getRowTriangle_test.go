package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRowTriangle(t *testing.T) {
	rows := []int{
		-1,
		0,
		2,
	}
	result := [][]int{
		nil,
		{1},
		{1, 2, 1},
	}
	for i, r := range rows {
		rr := GetRowTriangle(r)
		assert.Equal(t, rr, result[i])
	}
}
