package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// HasGroupsSizeX 卡牌分组 https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
func HasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}
	m := make(map[int]int, 0)
	for _, d := range deck {
		m[d]++
	}
	last := m[deck[0]]
	if last < 2 {
		return false
	}
	for _, count := range m {
		if count == last {
			continue
		}
		if last < count {
			last, count = count, last
		}
		last = utils.Gcd(last, count)
		if last < 2 {
			return false
		}
	}
	return true
}

