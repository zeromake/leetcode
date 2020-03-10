package math

import "math"

func Reverse(num int) int {
	var x int
	for num != 0 {
		x = x*10 + num%10
		if x > math.MaxInt32 || x < math.MinInt32 {
			return 0
		}
		num /= 10
	}
	return x
}
