package matrix

// SearchMatrix 搜索矩阵 https://leetcode-cn.com/problems/search-a-2d-matrix
func SearchMatrix(matrix [][]int, target int) bool {
	var m = len(matrix)
	if m == 0 {
		return false
	}
	var n = len(matrix[0])
	if n == 0 {
		return false
	}
	// 把矩阵转逻辑上换为有序数组，使用二分查找
	var (
		lens  = m * n
		left  = 0
		right = lens - 1
		mid   = (right + left) / 2
	)
	for left <= right {
		num := matrix[mid/n][mid%n]
		if num == target {
			return true
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (right + left) / 2
	}
	return false
}
