package matrix

func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			// [
			//[1,2,3],
			//[4,5,6],
			//[7,8,9],
			//]
			// 2, 0 : 7
			// 2, 1 : 8
			temp := matrix[n-1-j][i]
			// 2, 0: 7 = 2, 2: 9
			// 2, 1: 8 = 1, 2: 6
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			// 2, 2: 9 = 0, 2: 3
			// 1, 2: 6 = 0, 1: 2
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			// 0, 2: 3 = 0, 0: 1
			// 0, 1: 2 = 1, 0: 4
			matrix[j][n-1-i] = matrix[i][j]
			// 0, 0: 1 = 2, 0: 7
			// 2, 0: 7 = 2, 1: 8
			matrix[i][j] = temp
		}
	}
}
