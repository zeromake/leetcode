package arrays

// SortColors 颜色分类 https://leetcode-cn.com/problems/sort-colors
func SortColors(nums []int) {
	lens := len(nums)
	if lens < 2 {
		return
	}
	left, right := 0, lens
	i := 0
	for i < right {
		// 把数组切割为三段
		if nums[i] == 0 {
			// 把开头替换为 0 元素的数据
			swap(nums, i, left)
			left++
			i++
		} else if nums[i] == 1 {
			// 1 的直接等待即可
			i++
		} else {
			// 把结尾替换为 2 元素的数据
			right--
			swap(nums, i, right)
		}
	}
}

func swap(nums []int, index1, index2 int) {
	nums[index1], nums[index2] = nums[index2], nums[index1]
}
