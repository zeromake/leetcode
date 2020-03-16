package arrays

// RemoveDuplicates 删除排序数组中的重复项 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	offset := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != nums[offset] {
			if index != offset {
				nums[offset+1] = nums[index]
			}
			offset++
		}
	}
	return offset + 1
}
