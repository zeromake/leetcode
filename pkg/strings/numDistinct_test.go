package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	ss := [][2]string{
		{
			"rabbbit",
			"rabbit",
		},
		{
			"babgbag",
			"bag",
		},
	}
	result := []int{
		3,
		5,
	}
	for i, r := range ss {
		rr := NumDistinct(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
