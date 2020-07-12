package matrix

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// CalculateMinimumHP https://leetcode-cn.com/problems/dungeon-game
func CalculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			minn := utils.MinInt(dp[i+1][j], dp[i][j+1])
			dp[i][j] = utils.MaxInt(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}
