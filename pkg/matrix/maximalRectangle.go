package matrix

// MaximalRectangle 最大矩形 https://leetcode-cn.com/problems/maximal-rectangle
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	left, right, height := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	maxArea := 0
	for i := 0; i < m; i++ {
		curLeft, curRight := 0, n
		// height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// left
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if curLeft > left[j] {
					left[j] = curLeft
				}
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}
		// right
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if curRight < right[j] {
					right[j] = curRight
				}
			} else {
				right[j] = n
				curRight = j
			}
		}
		// area
		for j := 0; j < n; j++ {
			area := height[j] * (right[j] - left[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
