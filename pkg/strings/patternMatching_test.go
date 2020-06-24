package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatternMatching(t *testing.T) {
	strs := [][2]string{
		{"abba", "dogcatcatdog"},
		{"abba", "dogcatcatfish"},
		{"aaaa", "dogcatcatdog"},
		{"abba", "dogdogdogdog"},
		{"abb", "dogcatcat"},
		{"a", ""},
		{"", "cat"},
	}
	result := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
		false,
	}
	for i, r := range strs {
		rr := PatternMatching(r[0], r[1])
		assert.Equal(t, rr, result[i])
	}
}
