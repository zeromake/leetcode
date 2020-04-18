package arrays

// LargestRectangleArea 柱状图中的最大矩形 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func LargestRectangleArea(heights []int) int {
	l := len(heights)

	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	// 用切片模拟栈，为了节省切片扩容时间，预设切片空间为数组长度
	// 栈用来存heights数组中元素下标
	stack := make([]int, 1, l)
	stack[0] = -1 // 预填一个-1

	maxArea := 0
	var area int
	for i := 0; i < l; i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			area = heights[stack[len(stack)-1]] * (i - stack[len(stack)-2] - 1)
			if area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 { // 栈未到底
		area = heights[stack[len(stack)-1]] * (l - stack[len(stack)-2] - 1)
		if area > maxArea {
			maxArea = area
		}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}
