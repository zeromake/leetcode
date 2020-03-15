package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	nums := []int{
		3,
	}
	result := [][]string{
		{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		},
	}
	for i, r := range nums {
		rr := GenerateParenthesis(r)
		assert.Equal(t, rr, result[i])
	}
}
