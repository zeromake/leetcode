package matrix

// GameOfLife 生命游戏 https://leetcode-cn.com/problems/game-of-life
func GameOfLife(board [][]int) {
	var x, n = 0, len(board)
	if n == 0 {
		return
	}
	var m = len(board[0])
	// 8 个方向
	dir := [8][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x = 0
			for k := 0; k < 8; k++ {
				potion1, potion2 := i + dir[k][0], j + dir[k][1]
				if potion1 >= 0 && potion1 < n && potion2 >= 0 && potion2 < m && board[potion1][potion2] > 0 {
					x++
				}
			}
			if board[i][j] == 0 && x == 3 {
				// 用 -1 代表将要复活的标识
				board[i][j] = -1
			} else if board[i][j] == 1 && (x < 2 || x >3) {
				// 用 2 代表需要消亡的标识
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}

