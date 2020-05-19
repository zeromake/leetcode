package strings

// ValidPalindrome 验证回文字符串II https://leetcode-cn.com/problems/valid-palindrome-ii/
func ValidPalindrome(s string) bool {
	start, end := 0, len(s)-1
	// 通过左右指针检查非回文的位置
	for ; start < end && s[start] == s[end]; start, end = start+1, end-1 {
	}
	palindrome := func(i, j int) bool {
		for ; i < j && s[i] == s[j]; i, j = i+1, j-1 {
		}
		return i >= j
	}
	if start >= end {
		return true
	}
	// 检查两种可能
	return palindrome(start, end-1) || palindrome(start+1, end)
}
