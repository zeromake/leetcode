package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuperEggDrop(t *testing.T) {
	n := [][3]int{
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
	}
	for _, r := range n {
		rr := SuperEggDrop(r[0], r[1])
		//rr := SuperEggDrop2(r[0], r[1])
		assert.Equal(t, rr, r[2])
	}
}
