package arrays

// NumberOfSubarrays 统计优美子数组 https://leetcode-cn.com/problems/count-number-of-nice-subarrays
func NumberOfSubarrays(nums []int, k int) int {
	m := len(nums)
	counts := make([]int, m+1)
	odd, ans := 0, 0
	counts[0] = 1
	for i := 0; i < m; i++ {
		odd += nums[i] & 1
		if odd >= k {
			ans += counts[odd-k]
		}
		counts[odd]++
	}
	return ans
}
