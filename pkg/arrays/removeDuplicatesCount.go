package arrays

func RemoveDuplicatesCount(nums []int) int {
	offset := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[offset] = nums[i]
			offset++
		}
	}
	return offset
}
