package math

// SuperEggDrop 鸡蛋掉落 https://leetcode-cn.com/problems/super-egg-drop
func SuperEggDrop(k, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	count := 0
	for dp[k][count] < n {
		count++
		for i := 1; i <= k; i++ {
			dp[i][count] = dp[i][count-1] + dp[i-1][count-1] + 1
		}
	}
	return count
}
