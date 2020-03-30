package arrays

import "github.com/zeromake/leetcode/pkg/utils"

func Jump(nums []int) int {
	var (
		ans = 0
		end = 0
		maxPos = 0
	)
	for i := 0; i < len(nums) - 1; i++ {
		maxPos = utils.MaxInt(nums[i] + i, maxPos)
		if i == end {
			end = maxPos
			ans ++
		}
	}
	return ans
}

