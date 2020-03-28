package strings

// MinimumLengthEncoding 单词的压缩编码 https://leetcode-cn.com/problems/short-encoding-of-words
func MinimumLengthEncoding(words []string) int {
	cache := make(map[string]bool, len(words))
	for _, w := range words {
		cache[w] = true
	}
	for _, w := range words {
		for i := 1; i < len(w); i++ {
			_, ok := cache[w[i:]]
			if ok {
				delete(cache, w[i:])
			}
		}
	}
	count := 0
	for k := range cache {
		count += len(k) + 1
	}
	return count
}
