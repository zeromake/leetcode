package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	lines := [][4][]int{
		{
			{0, 0},
			{1, 0},
			{1, 1},
			{0, -1},
		},
		{
			{0, 0},
			{3, 3},
			{1, 1},
			{2, 2},
		},
		{
			{0, 0},
			{1, 1},
			{1, 0},
			{2, 1},
		},
		{
			{0, 0},
			{1, 1},
			{2, 2},
			{4, 4},
		},
		{
			{0, 0},
			{0, 5},
			{0, 1},
			{0, 2},
		},
	}
	result := [][]float64{
		{.5, 0},
		{1, 1},
		{},
		{},
		{0, 1},
	}
	for i, r := range lines {
		rr := Intersection(r[0], r[1], r[2], r[3])
		assert.Equal(t, rr, result[i])
	}
}
