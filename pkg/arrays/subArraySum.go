package arrays

// SubArraySum 和为K的子数组 https://leetcode-cn.com/problems/subarray-sum-equals-k/
func SubArraySum(nums []int, k int) int {
	count := 0
	pre := 0
	mp := make(map[int]int, 0)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := mp[pre-k]; ok {
			count += v
		}
		mp[pre]++
	}
	return count
}
