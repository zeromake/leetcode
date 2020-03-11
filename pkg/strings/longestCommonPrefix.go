package strings

// LongestCommonPrefix 最长公共前缀 https://leetcode-cn.com/problems/longest-common-prefix
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		count = 0
		// 另外一种先找出最短
		lens = len(strs[0])
		ss = strs[0]
		arr = strs[1:]
	)

loop:
	for i := 0; i < lens; i++ {
		c := ss[i]
		for _, str := range arr {
			// 长度不够或者无法匹配就结束
			if i >= len(str) || str[i] != c {
				break loop
			}
		}
		count++
	}
	return ss[0:count]
}
