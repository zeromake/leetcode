package strings

import "github.com/zeromake/leetcode/pkg/utils"

// LongestPalindrome 最长回文字符串 https://leetcode-cn.com/problems/longest-palindromic-substring
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var (
		start = 0
		end   = 0
	)
	for i := 0; i < len(s); i++ {
		// 扩展以 i 为中心
		len1 := expandAroundCenter(s, i, i)
		// 扩展以两个字符为中心
		len2 := expandAroundCenter(s, i, i+1)
		lens := utils.MaxInt(len1, len2)
		if lens > end-start {
			start = i - (lens-1)/2
			end = i + lens/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	var (
		lens = len(s)
	)
	// 直接向两边扩展
	for left >= 0 && right < lens && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
