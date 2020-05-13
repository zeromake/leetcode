package arrays

// GetRowTriangle 获取杨辉三角第n行 https://leetcode-cn.com/problems/pascals-triangle-ii
func GetRowTriangle(row int) []int {
	row++
	if row <= 0 {
		return nil
	}
	var (
		inner, help []int
		count       = 1
	)
	inner = []int{1}
	for row != count {
		help = inner
		inner = make([]int, count+1)
		inner[0] = 1
		inner[count] = 1
		for i := 1; i < count; i++ {
			inner[i] = help[i-1] + help[i]
		}
		count++
	}
	return inner
}
