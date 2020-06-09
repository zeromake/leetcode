package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslateNum(t *testing.T) {
	nums := []int{
		12258,
		25,
	}
	result := []int{
		5,
		2,
	}
	for i, r := range nums {
		rr := TranslateNum(r)
		assert.Equal(t, rr, result[i])
		rr = TranslateNum2(r)
		assert.Equal(t, rr, result[i])
		rr = TranslateNum3(r)
		assert.Equal(t, rr, result[i])
	}
}

func BenchmarkTranslateNum(b *testing.B) {
	b.ReportAllocs()
	num := 12258
	result := 5
	for i := 0; i < b.N; i++ {
		r := TranslateNum(num)
		if r != result {
			b.Error(r, result)
		}
	}
}
func BenchmarkTranslateNum2(b *testing.B) {
	b.ReportAllocs()
	num := 12258
	result := 5
	for i := 0; i < b.N; i++ {
		r := TranslateNum2(num)
		if r != result {
			b.Error(r, result)
		}
	}
}

func BenchmarkTranslateNum3(b *testing.B) {
	b.ReportAllocs()
	num := 12258
	result := 5
	for i := 0; i < b.N; i++ {
		r := TranslateNum3(num)
		if r != result {
			b.Error(r, result)
		}
	}
}
