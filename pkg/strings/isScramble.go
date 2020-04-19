package strings

// IsScramble 扰乱字符串 https://leetcode-cn.com/problems/scramble-string/
func IsScramble(s1, s2 string) bool {
	m, n := len(s1), len(s2)
	if m != n {
		return false
	}
	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for lens := 2; lens <= n; lens++ {
		for i := 0; i <= n-lens; i++ {
			for j := 0; j <= n-lens; j++ {
				for k := 1; k <= lens-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][lens-k] {
						dp[i][j][lens] = true
						break
					}
					if dp[i][j+lens-k][k] && dp[i+k][j][lens-k] {
						dp[i][j][lens] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

var tmp = map[byte]int{
	'a': 2,
	'b': 3,
	'c': 5,
	'd': 7,
	'e': 11,
	'f': 13,
	'g': 17,
	'h': 19,
	'i': 23,
	'j': 29,
	'k': 31,
	'l': 37,
	'm': 41,
	'n': 43,
	'o': 47,
	'p': 53,
	'q': 59,
	'r': 61,
	's': 67,
	't': 71,
	'u': 73,
	'v': 79,
	'w': 83,
	'x': 89,
	'y': 97,
	'z': 101,
}

func IsScramble2(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	length := len(s1)
	for i := 1; i < length; i++ {
		if mux(s1[:i]) == mux(s2[:i]) {
			if IsScramble2(s1[:i], s2[:i]) && IsScramble2(s1[i:], s2[i:]) {
				return true
			}

		}
		if mux(s1[:i]) == mux(s2[length-i:]) {
			if IsScramble2(s1[:i], s2[length-i:]) && IsScramble2(s1[i:], s2[0:length-i]) {
				return true
			}
		}
	}
	return false
}

func mux(s1 string) int {
	ans := 1
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		ans = ans * tmp[c]
	}
	return ans
}
