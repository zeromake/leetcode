package strings

// LengthOfLastWord 最后一个单词的长度 https://leetcode-cn.com/problems/length-of-last-word/
func LengthOfLastWord(s string) int {
	count := 0
	flag := false
	for i := len(s) - 1; i >=0; i-- {
		c := s[i]
		if c == ' ' {
			if flag {
				break
			} else {
				continue
			}
		} else if !flag {
			flag = true
		}
		count++
	}
	return count
}
