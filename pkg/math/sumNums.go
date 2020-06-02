package math

// SumNums 求1+2+…+n https://leetcode-cn.com/problems/qiu-12n-lcof/
func SumNums(n int) int {
	if n > 1 {
		n += SumNums(n - 1)
	}
	return n
}
