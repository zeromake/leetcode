package strings

// FindTheLongestSubstring 每个元音包含偶数次的最长子字符串 https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
func FindTheLongestSubstring(s string) int {
	positions := [32]int{}
	result, status := 0, 0
	for i := 0; i < 32; i++ {
		positions[i] = -1
	}
	positions[0] = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case 'a':
			status ^= 1
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		// 以位运算方式记录每次偶数个下的最新座标
		if ss := positions[status]; ss >= 0 {
			tmp := i - ss + 1
			if result < tmp {
				result = tmp
			}
		} else {
			positions[status] = i + 1
		}
	}
	return result
}
