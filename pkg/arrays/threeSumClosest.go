package arrays

import (
	"sort"

	"github.com/zeromake/leetcode/pkg/utils"
)

// ThreeSumClosest 最接近的三数之和 https://leetcode-cn.com/problems/3sum-closest/
func ThreeSumClosest(nums []int, target int) int {
	var (
		lens = len(nums)
	)
	if lens < 3 {
		return -1
	}
	closest := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < lens-2; i++ {
		left := i + 1
		right := lens - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			} else if sum > target {
				right--
			} else {
				left++
			}
			if utils.DistanceInt(sum, target) < utils.DistanceInt(closest, target) {
				closest = sum
			}
		}
	}
	return closest
}
