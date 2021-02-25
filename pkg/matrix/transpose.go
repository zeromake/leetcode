package matrix

// Transpose 转置矩阵 https://leetcode-cn.com/problems/transpose-matrix
func Transpose(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		row := matrix[i]
		for j := 0; j < m; j++ {
			result[j][i] = row[j]
		}
	}
	return result
}
