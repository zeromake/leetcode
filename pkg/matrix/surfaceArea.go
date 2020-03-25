package matrix

import "github.com/zeromake/leetcode/pkg/utils"

func SurfaceArea(grid [][]int) int {
	var (
		n, m = len(grid), len(grid[0])
		area = 0
	)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := grid[i][j]
			if count > 0 {
				area += count * 4 + 2
				if i > 0 {
					area -= utils.MinInt(grid[i - 1][j], count) * 2
				}
				if j > 0 {
					area -= utils.MinInt(grid[i][j - 1], count) * 2
				}
			}
		}
	}
	return area
}
