package matrix

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// UpdateMatrix 01 矩阵 https://leetcode-cn.com/problems/01-matrix
func UpdateMatrix(matrix [][]int) [][]int {
	var (
		m  = len(matrix)
		n  = len(matrix[0])
		dp = make([][]int, m)
	)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				dp[i][j] = math.MaxInt32 / 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 {
				dp[i][j] = utils.MinInt(dp[i][j], dp[i-1][j]+1)
			}
			if j >= 1 {
				dp[i][j] = utils.MinInt(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				dp[i][j] = utils.MinInt(dp[i][j], dp[i+1][j]+1)
			}
			if j+1 < n {
				dp[i][j] = utils.MinInt(dp[i][j], dp[i][j+1]+1)
			}
		}
	}
	return dp
}
