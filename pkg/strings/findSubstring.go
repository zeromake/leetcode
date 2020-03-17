package strings

func FindSubstring(s string, words []string) []int {
	var (
		result []int
	)
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	var (
		oneWord = len(words[0])
		wordNum = len(words)
	)
	m1 := make(map[string]int, 0)
	for _, w := range words {
		m1[w]++
	}
	// 外层为偏移
	for i := 0; i < oneWord; i++ {
		var (
			left = i
			right = i
			count = 0
		)
		m2 := make(map[string]int)
		for right + oneWord <= len(s) {
			w := s[right: right+oneWord]
			right += oneWord

			if _, ok := m1[w]; !ok {
				count = 0
				left = right
				m2 = map[string]int{}
			} else {
				m2[w] ++
				count ++
				// 出现重复情况清理最前面的子串
				for m2[w] > m1[w] {
					t := s[left: left + oneWord]
					count --
					m2[t]--
					left += oneWord
				}
				// 直到全部匹配
				if count == wordNum {
					result = append(result, left)
				}
			}
		}
	}
	return result
}

func FindSubstring2(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	wordLen := len(words[0])
	wordNum := len(words)
	sLen := len(s)
	res := make([]int, 0)
	l := wordLen * wordNum
	var left, cursor int
	reset := true
	right := l
	var m map[string]int8
	for i := 0; i < wordLen; i++ {
		left = i
		cursor = left
		right = left + l
		if reset {
			reset = false
			m = map[string]int8{}
			for _, w := range words {
				m[w]++
			}
		}

		for cursor < right && right <= sLen {
			index := s[cursor : cursor+wordLen]
			how, ok := m[index]
			if !ok {
				if reset {
					reset = false
					m = map[string]int8{}
					for _, w := range words {
						m[w]++
					}
				}
				left = cursor + wordLen
				cursor = left
				right = left + l
				continue
			}
			if how <= 0 {
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
				continue
			}
			reset = true
			m[index]--
			cursor += wordLen
			if cursor == right {
				res = append(res, left)
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
			}
		}
	}
	return res
}
