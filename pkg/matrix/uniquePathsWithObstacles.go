package matrix

// UniquePathsWithObstacles 不同路径 II https://leetcode-cn.com/problems/unique-paths-ii
func UniquePathsWithObstacles(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	// bottom
	for i := 0; i < m; i++ {
		// 只要一个出现障碍直接截断
		if grid[i][0] == 1 {
			break
		}
		cache[i][0] = 1
	}
	// right
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			break
		}
		cache[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 有障碍截断
			if grid[i][j] == 1 {
				cache[i][j] = 0
			} else {
				cache[i][j] = cache[i-1][j] + cache[i][j-1]
			}
		}
	}
	return cache[m-1][n-1]
}
