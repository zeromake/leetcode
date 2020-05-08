package matrix

import "github.com/zeromake/leetcode/pkg/utils"

// MaximalSquare 最大正方形 https://leetcode-cn.com/problems/maximal-square/
func MaximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	maxSide := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				max := utils.MinInt(utils.MinInt(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				dp[i][j] = max
				if max > maxSide {
					maxSide = max
				}
			}
		}
	}
	return maxSide * maxSide
}
