package matrix

const FULL = 'V'

// Solve 被围绕的区域 https://leetcode-cn.com/problems/surrounded-regions
func Solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	// 围绕着周围做 dfs
	for i := 0; i < n; i++ {
		solveDFS(0, i, m, n, board)
		solveDFS(m-1, i, m, n, board)
	}
	for i := 0; i < m; i++ {
		solveDFS(i, 0, m, n, board)
		solveDFS(i, n-1, m, n, board)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == FULL {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(i, j, m, n int, board [][]byte) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	} else if board[i][j] == 'X' || board[i][j] == FULL {
		// 已访问过或者不等于O
		return
	}
	board[i][j] = FULL
	solveDFS(i-1, j, m, n, board)
	solveDFS(i+1, j, m, n, board)
	solveDFS(i, j-1, m, n, board)
	solveDFS(i, j+1, m, n, board)
}
