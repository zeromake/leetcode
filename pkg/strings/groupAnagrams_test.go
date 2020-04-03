package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}
	result := []map[string][]string{
		{
			"eat": []string{"eat", "tea", "ate"},
			"tan": []string{"tan", "nat"},
			"bat": []string{"bat"},
		},
	}
	for i, r := range strs {
		rr := GroupAnagrams(r)
		res := result[i]
		for _, v := range rr {
			assert.Equal(t, v, res[v[0]])
		}
	}
}
