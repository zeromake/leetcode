package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// Rob 打家劫舍 https://leetcode-cn.com/problems/house-robber/
func Rob(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	if lens == 1 {
		return nums[0]
	}
	first, second := nums[0], utils.MaxInt(nums[0], nums[1])
	for i := 2; i < lens; i++ {
		first, second = second, utils.MaxInt(first+nums[i], second)
	}
	return second
}
