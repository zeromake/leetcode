package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZConvert(t *testing.T) {
	res := []string{
		"LEETCODEISHIRING",
		"LEETCODEISHIRING",
		"AB",
	}
	result := []string{
		"LCIRETOESIIGEDHN",
		"LDREOEIIECIHNTSG",
		"AB",
	}
	rows := []int{
		3,
		4,
		1,
	}
	for i, r := range res {
		rr := ZConvert(r, rows[i])
		assert.Equal(t, rr, result[i])
	}
}
