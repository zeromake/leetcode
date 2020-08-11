package strings

import "github.com/zeromake/leetcode/pkg/utils"

// CountBinarySubstrings 计数二进制子串 https://leetcode-cn.com/problems/count-binary-substrings/
func CountBinarySubstrings(s string) int {
	var ptr, last, ans int
	n := len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += utils.MinInt(count, last)
		last = count
	}
	return ans
}
