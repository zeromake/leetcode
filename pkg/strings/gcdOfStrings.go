package strings

func check(t, s string) bool {
	tLen := len(t)
	lens := len(s) / tLen
	for i := 0; i < lens; i += tLen {
		if s[i: i +tLen] != t {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a % b)
}

func GcdOfStrings(str1, str2 string) string {
	var (
		len1 = len(str1)
		len2 = len(str2)
	)
	// 最大公约数必然为最大公约字符串长度
	t := str1[:gcd(len1, len2)]
	// 检查是否为多个 t 拼接而成
	if check(t, str1) && check(t, str2) {
		return t
	}
	return ""
}

func GcdOfStrings2(str1, str2 string) string {
	// 相反顺序相加相同才能组成最大公约字符串
	if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}
