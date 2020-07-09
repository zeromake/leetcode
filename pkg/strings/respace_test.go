package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRespace(t *testing.T) {
	dictionary := [][]string{
		{"looked","just","like","her","brother"},
	}
	sentence := []string{
		"jesslookedjustliketimherbrother",
	}
	result := []int{
		7,
	}
	for i, r := range dictionary {
		rr := Respace(r, sentence[i])
		assert.Equal(t, rr, result[i])
	}
}
