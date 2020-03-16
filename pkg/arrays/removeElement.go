package arrays

// RemoveElement 删除元素 https://leetcode-cn.com/problems/remove-element
func RemoveElement(nums []int, val int) int {
	offset := 0
	for _, num := range nums {
		if num != val {
			nums[offset] = num
			offset++
		}
	}
	return offset
}
