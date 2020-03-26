package matrix

func NumRookCaptures(board [][]byte) int {
	var rook = [2]int{-1, -1}
	for i := 0; i < 8; i ++ {
		for j := 0; j < 8; j ++ {
			if board[i][j] == 'R' {
				rook[0] = i
				rook[1] = j
				break
			}
		}
	}
	acts := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	sum := 0
	for _, act := range acts {
		i, j := act[0] + rook[0], act[1] + rook[1]
		for i >= 0 && i < 8 && j >= 0 && j < 8 {
			aa := board[i][j]
			if aa == 'B' {
				break
			}
			if aa == 'p' {
				sum++
				break
			}
			i += act[0]
			j += act[1]
		}
	}
	return sum
}
