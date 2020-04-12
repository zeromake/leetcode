package strings

import "github.com/zeromake/leetcode/pkg/utils"

// MinDistance 编辑距离 https://leetcode-cn.com/problems/edit-distance/
func MinDistance(word1 string, word2 string) int {
	var (
		m, n = len(word1), len(word2)
	)
	if m == 0 || n == 0 {
		return m + n
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown++
			}
			dp[i][j] = utils.MinInt(left, utils.MinInt(down, leftDown))
		}
	}
	return dp[m][n]
}
