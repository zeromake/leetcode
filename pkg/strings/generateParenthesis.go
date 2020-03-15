package strings

// GenerateParenthesis 括号生成 https://leetcode-cn.com/problems/generate-parentheses
func GenerateParenthesis(n int) []string {
	var ans = &[]string{}
	backtrack(ans, "", 0, 0, n)
	return *ans
}

func backtrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	// 左括号未用完
	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	// 右括号可以使用
	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}
