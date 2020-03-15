package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	strs := []string{
		"()",
		"{}",
		"[]",
		"([)]",
		"{[]}",
		"{([{}])}",
		"?",
	}
	result := []bool{
		true,
		true,
		true,
		false,
		true,
		true,
		false,
	}
	for i, r := range strs {
		assert.Equal(t, IsValid(r), result[i])
	}
}
