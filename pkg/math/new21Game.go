package math

import "github.com/zeromake/leetcode/pkg/utils"

// New21Game 新21点游戏 https://leetcode-cn.com/problems/new-21-game/
func New21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W)
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1.0
	}

	dp[K-1] = 1.0 * float64(utils.MinInt(N-K+1, W)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+W+1]-dp[i+1])/float64(W)
	}
	return dp[0]
}
