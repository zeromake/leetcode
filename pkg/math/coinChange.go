package math

import "math"

// CoinChange 零钱兑换 https://leetcode-cn.com/problems/coin-change
func CoinChange(coins []int, amount int) int {
	var (
		lens = len(coins)
	)
	if lens == 0 {
		return -1
	}
	var cache = make([]int, amount+1)
	cache[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		// i 金额的最小硬币兑换数
		for _, coin := range coins {
			coin = i - coin
			if coin >= 0 && cache[coin] < min {
				min = cache[coin] + 1
			}
		}
		// 从小到大把所有的金额的最小硬币数保存起来
		cache[i] = min
	}
	// 如果没有匹配的返回 -1
	if cache[amount] == math.MaxInt32 {
		return -1
	}
	return cache[amount]
}

