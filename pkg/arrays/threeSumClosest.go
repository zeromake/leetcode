package arrays

import (
	"math"
	"sort"
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
	for i := 0; i < lens - 2; i ++ {
		left := i + 1
		right := lens - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			} else if sum > target {
				right --
			} else {
				left ++
			}
			if math.Abs(float64(sum - target)) < math.Abs(float64(closest - target)) {
				closest = sum
			}
		}
	}
	return closest
}
