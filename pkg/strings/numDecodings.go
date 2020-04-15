package strings

// NumDecodings 解码方法 https://leetcode-cn.com/problems/decode-ways
func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, curr := 1, 1 //dp[-1] = dp[0] = 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}
	return curr
}
