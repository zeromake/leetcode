package math

// LastRemaining 圆圈中最后剩下的数字 https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
func LastRemaining(n, m int) int {
	var (
		result = 0
	)
	for i := 2; i <= n; i++ {
		result = (m + result) % i
	}
	return result
}
