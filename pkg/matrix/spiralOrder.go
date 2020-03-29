package matrix

func SpiralOrder(matrix [][]int) []int {
	var (
		result []int
	)
	if len(matrix) == 0 {
		return result
	}
	var (
		left = 0
		top = 0
		right = len(matrix[0]) - 1
		bottom = len(matrix) - 1
		count = 0
		size = len(matrix) * len(matrix[0])
	)
	for left <= right {
		// top
		for i := left; i <= right; i++ {
			if count == size {
				break
			}
			result = append(result, matrix[top][i])
			count++
		}
		top++
		// right
		for i := top; i <= bottom; i++ {
			// 未找到测试用例
			//if count == size {
			//	break
			//}
			result = append(result, matrix[i][right])
			count++
		}
		right--
		// bottom
		for i := right; i >= left; i-- {
			if count == size {
				break
			}
			result = append(result, matrix[bottom][i])
			count++
		}
		bottom--
		// left
		for i := bottom ; i >= top; i -- {
			if count == size {
				break
			}
			result = append(result, matrix[i][left])
			count++
		}
		left++
	}
	return result
}
