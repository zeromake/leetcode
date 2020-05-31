package strings

import "unicode"

func isNumberOrChar(c rune) bool {
	return unicode.IsDigit(c) || unicode.IsLetter(c)
}

// IsPalindrome 验证回文串 https://leetcode-cn.com/problems/valid-palindrome/
func IsPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	low, high := 0, n-1
	for low < high {
		vi, vj := rune(s[low]), rune(s[high])
		if !isNumberOrChar(vi) && !isNumberOrChar(vj) {
			low++
			high--
		} else if !isNumberOrChar(vi) {
			low++
		} else if !isNumberOrChar(vj) {
			high--
		} else if unicode.ToLower(vi) != unicode.ToLower(vj) {
			return false
		} else {
			low++
			high--
		}
	}
	return true
}
