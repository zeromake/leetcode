package arrays

// GenerateTriangle 生成杨辉三角 https://leetcode-cn.com/problems/pascals-triangle/
func GenerateTriangle(row int) [][]int {
	if row == 0 {
		return nil
	}
	result := make([][]int, row)
	var (
		inner, help []int
		count       = 1
	)
	inner = []int{1}
	result[0] = inner
	for row != count {
		help = inner
		inner = make([]int, count+1)
		inner[0] = 1
		inner[count] = 1
		for i := 1; i < count; i++ {
			inner[i] = help[i-1] + help[i]
		}
		result[count] = inner
		count++
	}
	return result
}
