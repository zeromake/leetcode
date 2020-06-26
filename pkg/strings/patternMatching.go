package strings

// PatternMatching 模式匹配 https://leetcode-cn.com/problems/pattern-matching-lcci/
func PatternMatching(pattern string, value string) bool {
	m, n := 0, 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			m++
		} else {
			n++
		}
	}
	if m < n {
		m, n = n, m
		tmp := ""
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return n == 0
	}
	if len(pattern) == 0 {
		return false
	}

	for lens1 := 0; m*lens1 <= len(value); lens1++ {
		rest := len(value) - m*lens1
		if (n == 0 && rest == 0) || (n != 0 && rest%n == 0) {
			var lens2 int
			if n == 0 {
				lens2 = 0
			} else {
				lens2 = rest / n
			}
			pos, correct := 0, true
			var value1, value2 string
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == 'a' {
					sub := value[pos : pos+lens1]
					if len(value1) == 0 {
						value1 = sub
					} else if value1 != sub {
						correct = false
						break
					}
					pos += lens1
				} else {
					sub := value[pos : pos+lens2]
					if len(value2) == 0 {
						value2 = sub
					} else if value2 != sub {
						correct = false
						break
					}
					pos += lens2
				}
			}
			if correct && value1 != value2 {
				return true
			}
		}
	}
	return false
}
