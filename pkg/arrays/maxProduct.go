package arrays

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// MaxProduct 乘积最大子数组 https://leetcode-cn.com/problems/maximum-product-subarray/
func MaxProduct(nums []int) int {
	max := math.MinInt32
	iMax := 1
	iMin := 1
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = utils.MaxInt(iMax*val, val)
		iMin = utils.MinInt(iMin*val, val)
		max = utils.MaxInt(max, iMax)
	}
	return max
}
