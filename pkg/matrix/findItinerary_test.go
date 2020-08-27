package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tickets := [][][]string{
		{
			{"MUC", "LHR"},
			{"JFK", "MUC"},
			{"SFO", "SJC"},
			{"LHR", "SFO"},
		},
		{
			{"JFK", "SFO"},
			{"JFK", "ATL"},
			{"SFO", "ATL"},
			{"ATL", "JFK"},
			{"ATL", "SFO"},
		},
	}
	result := [][]string{
		{"JFK", "MUC", "LHR", "SFO", "SJC"},
		{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	}
	for i, r := range tickets {
		rr := FindItinerary(r)
		assert.Equal(t, rr, result[i])
	}
}
