package arrays

// LengthOfLIS 最长上升子序列 https://leetcode-cn.com/problems/longest-increasing-subsequence
func LengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	cell := []int{nums[0]}
	for _, num := range nums[1:] {
		lens := len(cell)
		// 发现比最后的数大说明是上升序列存到后面
		if num > cell[lens - 1] {
			cell = append(cell, num)
			continue
		}
		// 可优化为二分查找比当前小的数
		left := 0
		for i := lens - 1; i >= 0; i-- {
			if num > cell[i] {
				left = i + 1
				break
			}
		}
		cell[left] = num
	}
	return len(cell)
}

func LengthOfLIS2(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	cell := []int{nums[0]}
	for _, num := range nums[1:] {
		lens := len(cell)
		// 发现比最后的数大说明是上升序列存到后面
		if num > cell[lens - 1] {
			cell = append(cell, num)
			continue
		}
		// 可优化为二分查找比当前小的数
		left, right := 0, lens - 1
		for left < right {
			m := (left + right - 1) / 2
			if num > cell[m] {
				left = m + 1
			} else {
				right = m
			}
		}
		cell[left] = num
	}
	return len(cell)
}
