package strings

// IsMatchWildcard 通配符匹配 https://leetcode-cn.com/problems/wildcard-matching/
func IsMatchWildcard(s, p string) bool {
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, pLen + 1)
	}

	dp[0][0] = true
	for i := 1; i < pLen + 1; i++ {
		dp[0][i] = dp[0][i - 1] && p[i - 1] == '*'
	}

	for i := 1; i < sLen + 1; i++ {
		for j := 1; j < pLen + 1; j++ {
			if s[i - 1] == p[j - 1] || p[j - 1] == '?' {
				dp[i][j] = dp[i - 1][j - 1]
			} else if p[j - 1] == '*' {
				dp[i][j] = dp[i][j - 1] || dp[i - 1][j] || dp[i - 1][j - 1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[sLen][pLen]
}

func IsMatchWildcard2(s, p string) bool {
	i, j, m, n := 0, 0, 0, 0
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			j++
			n = j
			m = i
		} else if n > 0 {
			m++
			i = m
			j = n
		} else {
			return false
		}
	}
	// 字符全匹配后是否有多个 * 等待匹配
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}
