package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	s := []string{
		"test",
		"a ",
		" a",
		" a ",
		"  ",
	}
	result := []int{
		4,
		1,
		1,
		1,
		0,
	}
	for i, r := range s {
		rr := LengthOfLastWord(r)
		assert.Equal(t, rr, result[i])
	}
}
