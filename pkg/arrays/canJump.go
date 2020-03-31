package arrays

import "math"

func CanJump(nums []int) bool {
	var (
		ans = 0
		end = 0
		maxPos = 0
	)
	for i := 0; i < len(nums) - 1; i++ {
		maxPos = int(math.Max(float64(nums[i] + i), float64(maxPos)))
		if i == end {
			end = maxPos
			ans ++
		}
	}
	return end >= len(nums) - 1
}
