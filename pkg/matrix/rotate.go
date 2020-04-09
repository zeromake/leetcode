package matrix

// Rotate 旋转图像 https://leetcode-cn.com/problems/rotate-image 旋转矩阵 https://leetcode-cn.com/problems/rotate-matrix-lcci
func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			x, y := n-1-i, n-1-j
			// [
			//  [ 5, 1, 9,11],
			//  [ 2, 4, 8,10],
			//  [13, 3, 6, 7],
			//  [15,14,12,16]
			//]
			// 使用 4 x 4 例子
			// 取 [0, 0] 5 的旋转变换 [3, 0] 15, 变换规则为替换 y 为相反 x, y 互换
			temp := matrix[y][i]
			matrix[y][i] = matrix[x][y]
			matrix[x][y] = matrix[j][x]
			matrix[j][x] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}
