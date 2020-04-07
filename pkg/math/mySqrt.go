package math

import "math"

// MySort 平方根 https://leetcode-cn.com/problems/sqrtx
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	var (
		xx = float64(x)
		x0 = xx
		x1 = (x0 + xx / x0) / 2
	)
	for int(math.Abs(x0 - x1)) >= 1 {
		x0 = x1
		x1 = (x0 + xx / x0) / 2
	}
	return int(x1)
}

func MySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	//二分法
	left := 0
	right := x
	for left <= right {
		mid := left + (right-left) / 2
		temp := mid * mid
		if temp > x {
			right = mid - 1
		} else if temp < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left - 1
}
