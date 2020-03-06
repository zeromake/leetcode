package math

// 和为s的连续正数序列 https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
func FindContinuousSequence(target int) [][]int {
	var (
		result [][]int
		left   = 1
		right  = 2
	)
	// 发现 left 大于 target 说明已没有符合的可能性
	for left <= (target / 2) {
		// 等差数列计算公式
		sum := (left + right) * (right - left + 1) / 2
		// 匹配
		if sum == target {
			var res = make([]int, right-left+1)
			for i := left; i <= right; i++ {
				res[i-left] = i
			}
			result = append(result, res)
			// 把开始往右移动
			left++
		} else if sum < target {
			// 小于把结尾往右移动
			right++
		} else {
			// 大于移动头
			left++
		}
	}
	return result
}
