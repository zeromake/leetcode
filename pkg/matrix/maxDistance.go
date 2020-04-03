package matrix

// MaxDistance 地图分析 https://leetcode-cn.com/problems/as-far-from-land-as-possible/
func MaxDistance(grid [][]int) int {
	var (
		n = len(grid)
	)
	if n == 0 {
		return -1
	}
	m := len(grid[0])
	var group [][2]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			v := grid[i][j]
			if v == 1 {
				group = append(group, [2]int{i, j})
			}
		}
	}
	if len(group) == 0 || len(group) == n*m {
		return -1
	}
	count := 0
	for len(group) > 0 {
		temp := group
		group = nil
		// 每一大轮填海距离加一
		count++
		// 所有陆地一起填海
		for _, point := range temp {
			for _, dir := range dirs {
				x, y := point[0]+dir[0], point[1]+dir[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
					grid[x][y] = 1
					group = append(group, [2]int{x, y})
				}
			}
		}
	}
	return count - 1
}
