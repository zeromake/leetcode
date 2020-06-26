package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// MaxScoreSightseeingPair 最佳观光组合 https://leetcode-cn.com/problems/best-sightseeing-pair/
func MaxScoreSightseeingPair(A []int) int {
	ans, mx := 0, A[0]+0
	for j := 1; j < len(A); j++ {
		ans = utils.MaxInt(ans, mx+A[j]-j)
		// 边遍历边维护
		mx = utils.MaxInt(mx, A[j]+j)
	}
	return ans
}
