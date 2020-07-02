package matrix

// KthSmallest 有序矩阵中第K小的元素 https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
func KthSmallest(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	left, right := matrix[0][0], matrix[rows-1][cols-1]
	// 由于可能会出现重复，所以采用这种方式的二分
	for left < right {
		mid := (left + right) / 2
		num := count(matrix, mid, rows, cols)
		if num >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func count(matrix [][]int, target int, rows, cols int) int {
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] <= target {
				ans++
			} else {
				break
			}
		}
	}
	return ans
}
