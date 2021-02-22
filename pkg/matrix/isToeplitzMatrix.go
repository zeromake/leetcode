package matrix

// IsToeplitzMatrix 托普利茨矩阵 https://leetcode-cn.com/problems/toeplitz-matrix/
func IsToeplitzMatrix(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	if n == 1 && m == 1 {
		return true
	}
	var (
		y           = 0
		rowCount    = n - 1
		columnCount = m - 1
	)
	// 以行去遍历矩阵，用当前行和下一行去判断，最后一行和最后一列不用处理。
	for y < rowCount {
		x := 0
		for x < columnCount {
			if matrix[y][x] != matrix[y+1][x+1] {
				return false
			}
			x++
		}
		y++
	}
	return true
}
