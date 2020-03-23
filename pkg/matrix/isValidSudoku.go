package matrix

// IsValidSudoku 有效数独 https://leetcode-cn.com/problems/valid-sudoku
func IsValidSudoku(board [][]byte) bool {
	var (
		row, col = [9][9]bool{}, [9][9]bool{}
		block = [9][9]bool{}
	)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				num := c - '1'
				index := i / 3 * 3 + j / 3
				if row[i][num] || col[j][num] || block[index][num] {
					return false
				} else {
					row[i][num] = true
					col[j][num] = true
					block[index][num] = true
				}
			}
		}
	}
	return true
}
