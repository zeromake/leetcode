package strings

// StrStr 查找子串 https://leetcode-cn.com/problems/implement-strstr
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	offset := 0
	for i := 0; i < len(haystack); i++ {
		b := haystack[i]
		if b == needle[offset] {
			offset ++
		} else {
			i -= offset
			offset = 0
		}
		if offset == len(needle) {
			return i - offset + 1
		}
	}
	return -1
}
