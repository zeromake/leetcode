package strings

// ReverseWordsIII 反转字符串中的单词 III https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
func ReverseWordsIII(s string) string {
	var sb = []byte(s)
	var start int
	for i, b := range sb {
		if b == ' ' {
			reverseBytes(sb, start, i - 1)
			start = i + 1
		}
	}
	if start < len(sb) {
		reverseBytes(sb, start, len(sb) - 1)
	}
	return string(sb)
}

func reverseBytes(sb []byte, start, end int) {
	mid := start + (end-start)/2
	for i := start; i <= mid; i++ {
		var offset = i - start
		sb[i], sb[end-offset] = sb[end-offset], sb[i]
	}
}
