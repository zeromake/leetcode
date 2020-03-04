package matrix

import (
	"container/list"
)

// orangesRotting 检查腐烂橘子需要多少轮 https://leetcode-cn.com/problems/rotting-oranges
func orangesRotting(grid [][]int) int {
	var (
		M     = len(grid)
		N     = len(grid[0])
		queue = list.New()
		// 新鲜橘子数量
		count = 0
		// 腐烂分种数
		round = 0
	)
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			o := grid[r][c]
			if o == 1 {
				count++
			} else if o == 2 {
				queue.PushBack([2]int{r, c})
			}
		}
	}
	for count > 0 && queue.Len() != 0 {
		round++
		var n = queue.Len()
		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			var (
				orange = e.Value.([2]int)
				r      = orange[0]
				c      = orange[1]
			)
			// 向上检查
			if r-1 >= 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				count--
				queue.PushBack([2]int{r - 1, c})
			}
			// 向下检查
			if r+1 < M && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				count--
				queue.PushBack([2]int{r + 1, c})
			}
			// 向左检查
			if c-1 >= 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				count--
				queue.PushBack([2]int{r, c - 1})
			}
			// 向右检查
			if c+1 < N && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				count--
				queue.PushBack([2]int{r, c + 1})
			}
		}
	}
	if count > 0 {
		return -1
	} else {
		return round
	}
}

var (
	// 正好对应每个方向的座标变换
	dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

func orangesRotting2(grid [][]int) int {
	var (
		N     = len(grid)
		M     = len(grid[0])
		queue = list.New()
		// 新鲜橘子数量
		count = 0
		// 腐烂分种数
		round = 0
	)
	for r := 0; r < N; r++ {
		for c := 0; c < M; c++ {
			if grid[r][c] == 1 {
				count++
			} else if grid[r][c] == 2 {
				queue.PushBack([2]int{r, c})
			}
		}
	}
	for count > 0 && queue.Len() != 0 {
		round++
		var n = queue.Len()
		for i := 0; i < n; i++ {
			// 取出头部进行 bfs(广度优先遍历)
			e := queue.Front()
			queue.Remove(e)
			var (
				orange = e.Value.([2]int)
			)
			// 把四次判断转为循环
			for _, cur := range dirs {
				var (
					x = orange[0] + cur[0]
					y = orange[1] + cur[1]
				)
				if x >= 0 && x < N && y >= 0 && y < M && grid[x][y] == 1 {
					// 设置为腐烂
					grid[x][y] = 2
					// 减少新鲜的橘子
					count--
					// 把腐烂的橘子加入队列
					queue.PushBack([2]int{x, y})
				}
			}
		}
	}
	if count > 0 {
		return -1
	} else {
		return round
	}
}
