package strings

func IsMatchWildcard(s, p string) bool {
	var (
		n = len(s)
		m = len(p)
		dp = make([][]bool, n + 1)
	)
	for i, _ := range dp {
		dp[i] = make([]bool, m + 1)
	}
	dp[0][0] = true
	for i, c := range p {
		if c == '*' && dp[0][i] {
			dp[0][i+1] = true
		}
	}
	for i := 1; i <= n; i++ {
		a := s[i - 1]
		for j := 1; j <= m; j++ {
			b := p[j - 1]
			if a == b || b == '.' {
				dp[i][j] = dp[i - 1][j - 1]
			} else if b == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[n][m]
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
