package strings

import "strconv"

// CompressString 字符串压缩 https://leetcode-cn.com/problems/compress-string-lcci
func CompressString(S string) string {
	if len(S) == 0 {
		return S
	}
	cache := make([]byte, 0)
	prev := S[0]
	count := 1
	for i := 1; i < len(S); i++ {
		if prev != S[i] {
			cache = append(cache, prev)
			cache = append(cache, strconv.Itoa(count)...)
			count = 1
			prev = S[i]
		} else {
			count += 1
		}
	}
	cache = append(cache, prev)
	cache = append(cache, strconv.Itoa(count)...)
	if len(cache) >= len(S) {
		return S
	}
	return string(cache)
}
