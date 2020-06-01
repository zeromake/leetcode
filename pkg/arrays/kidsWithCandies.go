package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// KidsWithCandies 拥有最多糖果的孩子 https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/submissions/
func KidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for i := 0; i < n; i++ {
		maxCandies = utils.MaxInt(maxCandies, candies[i])
	}
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = candies[i]+extraCandies >= maxCandies
	}
	return ret
}
