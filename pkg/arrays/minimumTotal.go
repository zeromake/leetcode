package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// MinimumTotal 三角形最小路径和 https://leetcode-cn.com/problems/triangle/
func MinimumTotal(triangle [][]int) int {
	n := len(triangle) - 2
	for i := n; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += utils.MinInt(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
