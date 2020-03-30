package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatchWildcard(t *testing.T) {
	strs := [][2]string{
		{
			"aa",
			"a",
		},
		{
			"aa",
			"*",
		},
		{
			"cb",
			"?a",
		},
		{
			"a",
			"a***",
		},
	}
	result := []bool{
		false,
		true,
		false,
		true,
	}
	for i, r := range strs {
		rr := IsMatchWildcard(r[0], r[1])
		assert.Equal(t, rr, result[i])
		rr = IsMatchWildcard2(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
