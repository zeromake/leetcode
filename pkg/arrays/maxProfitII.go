package arrays

// MaxProfitII 买卖股票的最佳时机 II https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func MaxProfitII(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i -1] {
			result += prices[i] - prices[i -1]
		}
	}
	return result
}
