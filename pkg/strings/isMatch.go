package strings

func IsMatch(s, p string) bool {
	if p != "" && p[0] == '*' {
		return false
	}
	dp := make([][]bool, len(s) + 1)
	lens := len(p)
	for i, _ := range dp {
		dp[i] = make([]bool, lens + 1)
	}
	dp[0][0] = true
	for i, c := range p {
		if c == '*' && dp[0][i-1] {
			dp[0][i + 1] = true
		}
	}
	for i, a := range s {
		for j, b := range p {
			if a == b || b == '.' {
				dp[i + 1][j + 1] = dp[i][j]
			} else if b == '*' {
				if p[j - 1] != s[i] && p[j - 1] != '.' {
					dp[i + 1][j + 1]= dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
			}
		}
	}
	return dp[len(s)][lens]
}


func IsMatch2(s, p string) bool {
	if p == "" {
		return s == p
	}
	match := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return IsMatch2(s, p[2:]) || (match && IsMatch2(s[1:], p))
	} else {
		return match && IsMatch2(s[1:], p[1:])
	}
}
