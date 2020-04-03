package matrix

// generateMatrix 生成螺旋矩阵 https://leetcode-cn.com/problems/spiral-matrix-ii
func GenerateSpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, top, right, bottom := 0, 0, n-1, n-1
	num := 1
	// 由于是正方形判断左右即可
	for left <= right {
		// top
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// right
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// bottom
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// left
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
