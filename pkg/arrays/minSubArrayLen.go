package arrays

import (
	"math"

	"github.com/zeromake/leetcode/pkg/utils"
)

// MinSubArrayLen 长度最小的子数组 https://leetcode-cn.com/problems/minimum-size-subarray-sum/
func MinSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = utils.MinInt(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
