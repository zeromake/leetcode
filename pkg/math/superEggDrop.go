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

//func SuperEggDrop2(K, N int) int {
//	memo := make([][]int, K+1)
//	for i := 0; i <= K; i++ {
//		memo[i] = make([]int, N+1)
//	}
//	var dp func(k, n int) int
//	dp = func(k, n int) int {
//		if k == 1 {
//			return n
//		}
//		if n <= 0 {
//			return 0
//		}
//		if memo[k][n] != 0 {
//			return memo[k][n]
//		}
//		tmp := math.MaxInt32
//		var left, right = 1, n
//		for left <= right {
//			mid := (left + right) / 2
//			eggBreak := dp(k-1, mid-1)
//			eggUnBreak := dp(k, n-mid)
//			if eggBreak > eggUnBreak {
//				right = mid - 1
//				tmp = utils.MinInt(tmp, eggBreak+1)
//			} else {
//				left = mid + 1
//				tmp = utils.MinInt(tmp, eggUnBreak+1)
//			}
//		}
//		memo[k][n] = tmp
//		return tmp
//	}
//	return dp(K, N)
//}
