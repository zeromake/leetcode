package arrays

// MaxSubArray 最大和字串 https://leetcode-cn.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	var (
		max = nums[0]
		sum = 0
	)
	for _, num := range nums {
		if sum > 0 {
			// 只要 sum 还不为 <= 0 就加上连续的下一个数
			sum += num
		} else {
			// 负数开始和正负相抵的情况，需要重新开始
			sum = num
		}
		if max < sum {
			// 记录每次相加的最大数
			max = sum
		}
	}
	return max
}
