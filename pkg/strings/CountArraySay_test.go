package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountArraySay(t *testing.T) {
	result := []string{
		"1",
		"11",
		"21",
		"1211",
		"111221",
		"312211",
		"13112221",
		"1113213211",
		"111221",
		"312211",
	}
	n := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		5,
		6,
	}
	for i, r := range result {
		rr := CountArraySay(n[i])
		assert.Equal(t, rr, r)
	}
}
