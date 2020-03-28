package matrix

// SolveSudoku 解数独 https://leetcode-cn.com/problems/sudoku-solver
func SolveSudoku(board [][]byte) {
	var (
		row = [9][9]bool{}
		col = [9][9]bool{}
		box = [9][9]bool{}
	)

	// 遍历所有把已填入的记录下来
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				num := c - '1'
				index := i/3*3 + j/3
				row[i][num] = true
				col[j][num] = true
				box[index][num] = true
			}
		}
	}
	recusiveSolveSudoku(board, row, col, box, 0, 0)
}

func recusiveSolveSudoku(board [][]byte, row, col, box [9][9]bool, i, j int) bool {
	if j == 9 {
		j = 0
		i++
		if i == 9 {
			return true
		}
	}
	if board[i][j] == '.' {
		for num := 0; num < 9; num++ {
			index := i/3*3 + j/3
			use := !(row[i][num] || col[j][num] || box[index][num])
			// 检查是否可用
			if use {
				row[i][num] = true
				col[j][num] = true
				box[index][num] = true
				board[i][j] = byte('1' + num)
				// 设置上并递归判断
				if recusiveSolveSudoku(board, row, col, box, i, j+1) {
					return true
				}
				// 不能够填入，撤销填入
				board[i][j] = '.'
				row[i][num] = false
				col[j][num] = false
				box[index][num] = false
			}
		}
	} else {
		return recusiveSolveSudoku(board, row, col, box, i, j+1)
	}
	return false
}

func SolveSudoku2(board [][]byte) {
	var (
		row    = [9][9]bool{}
		col    = [9][9]bool{}
		box    = [9][9]bool{}
		whites [][2]int
	)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				num := c - '1'
				index := i/3*3 + j/3
				row[i][num] = true
				col[j][num] = true
				box[index][num] = true
			} else {
				// 收集需要填入的座标
				whites = append(whites, [2]int{i, j})
			}
		}
	}
	recusiveSolveSudoku2(board, row, col, box, whites, 0)
}

func recusiveSolveSudoku2(board [][]byte, row, col, box [9][9]bool, whites [][2]int, index int) bool {
	if index == len(whites) {
		return true
	}
	i, j := whites[index][0], whites[index][1]
	for num := 0; num < 9; num++ {
		boxIndex := i/3*3 + j/3
		use := !(row[i][num] || col[j][num] || box[boxIndex][num])
		if use {
			row[i][num] = true
			col[j][num] = true
			box[boxIndex][num] = true
			board[i][j] = byte('1' + num)
			if recusiveSolveSudoku2(board, row, col, box, whites, index+1) {
				return true
			}
			board[i][j] = '.'
			row[i][num] = false
			col[j][num] = false
			box[boxIndex][num] = false
		}
	}
	return false
}
