package matrix

import "container/list"

// NumIslands 岛屿数量 https://leetcode-cn.com/problems/number-of-islands
func NumIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	queue := list.New()
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				grid[i][j] = '0'
				queue.PushBack([2]int{i, j})
				for queue.Len() > 0 {
					e := queue.Front()
					queue.Remove(e)
					element := e.Value.([2]int)
					for _, dir := range dirs {
						x, y := element[0]+dir[0], element[1]+dir[1]
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
							queue.PushBack([2]int{x, y})
							grid[x][y] = '0'
						}
					}
				}
			}
		}
	}
	return result
}
