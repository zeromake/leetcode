package math

import "github.com/zeromake/leetcode/pkg/utils"

func MaxArea(heights []int) int {
	var (
		lens  = len(heights)
		start = 0
		end   = lens - 1
		sum   = 0
	)
	for start < end {
		width := end - start
		// 哪边的高度低就往中间走
		if heights[start] < heights[end] {
			sum = utils.MaxInt(sum, heights[start]*width)
			start++
		} else {
			sum = utils.MaxInt(sum, heights[end]*width)
			end--
		}
	}
	return sum
}
