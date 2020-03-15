package matrix

import (
	"container/list"
)

func MaxAreaOfIsland(grid [][]int) int {
	var (
		height  = len(grid)
		width   = len(grid[0])
		stack   = list.New()
		maxArea = 0
	)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 0 {
				continue
			}
			stack.PushBack([]int{i, j})
			area := 0
			for stack.Len() != 0 {
				e := stack.Front()
				stack.Remove(e)
				item := e.Value.([]int)
				x, y := item[0], item[1]
				if x < 0 || y < 0 || x >= height || y >= width || grid[x][y] == 0 {
					continue
				}
				area++
				grid[x][y] = 0
				for _, d := range dirs {
					stack.PushBack([]int{
						x + d[0],
						y + d[1],
					})
				}
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func MaxAreaOfIsland2(grid [][]int) int {
	var (
		height  = len(grid)
		width   = len(grid[0])
		maxArea = 0
		area    = 0
		dfs     func(x, y int)
	)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= height || y >= width || grid[x][y] == 0 {
			return
		}
		area += 1
		grid[x][y] = 0
		for _, d := range dirs {
			dfs(
				x+d[0],
				y+d[1],
			)
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			area = 0
			if grid[i][j] == 1 {
				dfs(i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
