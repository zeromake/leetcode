package strings

// NumDistinct 不同的子序列 https://leetcode-cn.com/problems/distinct-subsequences/
func NumDistinct(s, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n]
}
