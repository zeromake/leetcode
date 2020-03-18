package strings

// LongestValidParentheses 最大连续有效括号 https://leetcode-cn.com/problems/longest-valid-parentheses
func LongestValidParentheses(s string) int {
	var (
		left, right = 0, 0
		max = 0
	)
	// 正向计数左右括号，右大于左清零
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if max < right * 2 {
				max = right * 2
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	// 反向向计数左右括号，左大于右清零
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if max < right * 2 {
				max = right * 2
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return max
}

