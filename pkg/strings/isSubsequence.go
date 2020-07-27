package strings

// IsSubsequence 判断子序列 https://leetcode-cn.com/problems/is-subsequence/submissions/
func IsSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 {
		return true
	}
	if n > m {
		return false
	}
	j := 0
	for i := range t {
		if s[j] == t[i] {
			j++
		}
		if j == n {
			return true
		}
	}
	return j == n
}
