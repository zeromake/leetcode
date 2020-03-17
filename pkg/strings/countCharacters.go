package strings

func CountCharacters(words []string, chars string) int {
	var (
		count = 0
		cache [26]int
	)
	for i := 0; i < len(chars); i++ {
		cache[chars[i] - 'a'] += 1
	}
	for _, word := range words {
		// 直接赋值 map 会浅拷贝
		c, match := cache, true
		for i := 0; i < len(word); i++ {
			if c[word[i]- 'a'] <= 0 {
				match = false
				break
			}
			c[word[i] - 'a'] -= 1
		}
		if match {
			count += len(word)
		}
	}
	return count
}

