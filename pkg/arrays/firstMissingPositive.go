package arrays

// FirstMissingPositive 缺失的第一个正数 https://leetcode-cn.com/problems/first-missing-positive
func FirstMissingPositive(nums []int) int {
	var (
		lens = len(nums)
	)
	for i := 0; i < lens; i++ {
		for nums[i] > 0 && nums[i] <= lens && nums[nums[i]-1] != nums[i] {
			num := nums[i]
			nums[i], nums[num-1] = nums[num-1], nums[i]
		}
	}
	for i := 0; i < lens; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return lens + 1
}
