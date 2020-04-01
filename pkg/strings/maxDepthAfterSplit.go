package strings

// MaxDepthAfterSplit 有效括号的嵌套深度 https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
func MaxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))
	for i := 0; i < len(seq); i ++ {
		c := seq[i]
		if c == '(' {
			ans[i] = i & 1
		} else {
			ans[i] = (i + 1) & 1
		}
	}
	return ans
}
