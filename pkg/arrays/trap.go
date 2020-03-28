package arrays

// Trap 接雨水 https://leetcode-cn.com/problems/trapping-rain-water
func Trap(height []int) int {
	var (
		left     = 0
		right    = len(height) - 1
		leftMax  = 0
		rightMax = 0
		area     = 0
	)
	for left < right {
		// 双指针扫描
		if height[left] < height[right] {
			// 只有找到比上次的最大高度低才能存的了水
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				area += leftMax - height[left]
			}
			left++
		} else {
			// 只有找到比上次的最大高度低才能存的了水
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				area += rightMax - height[right]
			}
			right--
		}
	}
	return area
}
