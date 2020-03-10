package arrays

import "math"

// 买卖股票的最佳时机  https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
func MaxProfit(prices []int) int {
	var (
		sum = 0
		min = math.MaxInt32
	)
	for _, price := range prices {
		// 由于卖出日一定大于买入所以可以先取最小值再算盈利
		if price < min {
			min = price
		} else if price-min > sum {
			sum = price - min
		}
	}
	return sum
}
