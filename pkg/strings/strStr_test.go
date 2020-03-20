package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	strs := [][2]string{
		{"hello", "ll"},
		{"hello", ""},
		{"hh", "hello"},
	}
	result := []int{
		2,
		0,
		-1,
	}
	for i, r := range strs {
		rr := StrStr(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
