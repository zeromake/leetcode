package matrix

import "github.com/zeromake/leetcode/pkg/utils"

// https://leetcode-cn.com/problems/minimum-path-sum
func MinPathSum(grid [][]int) int {
	var (
		m, n = len(grid) - 1, len(grid[0]) - 1
	)
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if i != 0 && j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else if i != 0 && j != 0 {
				grid[i][j] += utils.MinInt(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m][n]
}
