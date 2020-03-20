package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := [][]int{
		{
			1, 2, 3, 6,
		},
		{
			4, 4, 5, 5, 5,
		},
		{},
		{
			1,
		},
		{
			1,
		},
		{
			4, 4, 4, 4, 5, 5, 5,
		},
		{
			4, 4, 5, 5, 5,
		},
	}
	target := []int{
		3,
		5,
		1,
		1,
		0,
		4,
		4,
	}
	result := [][]int{
		{2, 2},
		{2, 4},
		{-1, -1},
		{0, 0},
		{-1, -1},
		{0, 3},
		{0, 1},
	}
	for i, r := range nums {
		rr := SearchRange(r, target[i])
		assert.Equal(t, rr, result[i])
		rr = SearchRange2(r, target[i])
		assert.Equal(t, rr, result[i])
		rr = SearchRange3(r, target[i])
		assert.Equal(t, rr, result[i])
		rr = SearchRange4(r, target[i])
		assert.Equal(t, rr, result[i])
	}
}

func BenchmarkSearchRange(b *testing.B) {
	b.ReportAllocs()
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 11, 12, 13, 14, 15, 16, 17}
	target := 5
	result := [2]int{
		4,
		6,
	}
	for i := 0; i < b.N; i++ {
		rr := SearchRange(nums, target)
		if rr[0] != result[0] || rr[1] != result[1] {
			b.Error(rr, result)
		}
	}
}

func BenchmarkSearchRange2(b *testing.B) {
	b.ReportAllocs()
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 11, 12, 13, 14, 15, 16, 17}
	target := 5
	result := [2]int{
		4,
		6,
	}
	for i := 0; i < b.N; i++ {
		rr := SearchRange2(nums, target)
		if rr[0] != result[0] || rr[1] != result[1] {
			b.Error(rr, result)
		}
	}
}

func BenchmarkSearchRange3(b *testing.B) {
	b.ReportAllocs()
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 11, 12, 13, 14, 15, 16, 17}
	target := 5
	result := [2]int{
		4,
		6,
	}
	for i := 0; i < b.N; i++ {
		rr := SearchRange3(nums, target)
		if rr[0] != result[0] || rr[1] != result[1] {
			b.Error(rr, result)
		}
	}
}

func BenchmarkSearchRange4(b *testing.B) {
	b.ReportAllocs()
	nums := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 11, 12, 13, 14, 15, 16, 17}
	target := 5
	result := [2]int{
		4,
		6,
	}
	for i := 0; i < b.N; i++ {
		rr := SearchRange4(nums, target)
		if rr[0] != result[0] || rr[1] != result[1] {
			b.Error(rr, result)
		}
	}
}
