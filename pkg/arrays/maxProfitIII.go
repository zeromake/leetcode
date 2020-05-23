package arrays

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// MaxProfitIII 买股票的最佳时机III https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
func MaxProfitIII(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit1 := 0
	maxProfitAfterBuy := math.MinInt32
	maxProfit2 := 0
	for _, price := range prices {
		// 第一次最小购买的价格
		minPrice = utils.MinInt(minPrice, price)
		// 第一次卖出的最大利润
		maxProfit1 = utils.MaxInt(maxProfit1, price-minPrice)
		// 第二次购买后的剩余利润
		maxProfitAfterBuy = utils.MaxInt(maxProfitAfterBuy, maxProfit1-price)
		// 第二次卖出后的总利润
		maxProfit2 = utils.MaxInt(maxProfit2, price+maxProfitAfterBuy)
	}
	return maxProfit2
}
