package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquationsPossible(t *testing.T) {
	strs := [][]string{
		{"a==b", "b!=a"},
		{"b==a", "a==b"},
		{"a==b", "b==c", "a==c"},
		{"a==b", "b!=c", "a==c"},
		{"c==c", "b==d", "x!=z"},
	}
	result := []bool{
		false,
		true,
		true,
		false,
		true,
	}
	for i, r := range strs {
		rr := EquationsPossible(r)
		assert.Equal(t, rr, result[i])
	}
}
