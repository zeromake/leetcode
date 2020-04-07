package math

// ClimbStairs 爬楼梯 https://leetcode-cn.com/problems/climbing-stairs
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a + b
	}
	return b
}
