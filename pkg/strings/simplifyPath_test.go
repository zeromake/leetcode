package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	paths := []string{
		"/../",
		"/home/../etc/",
		"/root/./.git",
	}
	result := []string{
		"/",
		"/etc",
		"/root/.git",
	}
	for i, r := range paths {
		rr := SimplifyPath(r)
		assert.Equal(t, rr, result[i])
	}
}

