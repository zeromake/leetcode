package matrix

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

// UpdateBoard 扫雷游戏 https://leetcode-cn.com/problems/minesweeper/
func UpdateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		updateBoardBfs(board, x, y)
	}
	return board
}

func updateBoardBfs(board [][]byte, sx, sy int) {
	var queue [][]int
	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, len(board[0]))
	}

	queue = append(queue, []int{sx, sy})
	vis[sx][sy] = true
	for i := 0; i < len(queue); i++ {
		cnt, x, y := 0, queue[i][0], queue[i][1]
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
				continue
			}
			// 不用判断 M，因为如果有 M 的话游戏已经结束了
			if board[tx][ty] == 'M' {
				cnt++
			}
		}
		if cnt > 0 {
			board[x][y] = byte(cnt + '0')
		} else {
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				tx, ty := x+dirX[i], y+dirY[i]
				// 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
				if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' || vis[tx][ty] {
					continue
				}
				queue = append(queue, []int{tx, ty})
				vis[tx][ty] = true
			}
		}
	}
}
