package arrays

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// MinCostTickets 最低票价 https://leetcode-cn.com/problems/minimum-cost-for-tickets/
func MinCostTickets(days []int, costs []int) int {
	memo := [366]int{}
	durations := []int{1, 7, 30}

	var dp func(idx int) int
	dp = func(idx int) int {
		if idx >= len(days) {
			return 0
		}
		if memo[idx] > 0 {
			return memo[idx]
		}
		memo[idx] = math.MaxInt32
		j := idx
		for i := 0; i < 3; i++ {
			for ; j < len(days) && days[j] < days[idx]+durations[i]; j++ {
			}
			memo[idx] = utils.MinInt(memo[idx], dp(j)+costs[i])
		}
		return memo[idx]
	}
	return dp(0)
}
