package strings

import (
	"strconv"
	"strings"
)

var (
	dp = []string{"1"}
)

// CountArraySay 外观数列 https://leetcode-cn.com/problems/count-and-say/
func CountArraySay(n int) string {
	sb := strings.Builder{}
	i := 1
	if n < len(dp) {
		return dp[n-1]
	}
	for ; i < n; i++ {
		count := 1
		sb.Reset()
		last := dp[i-1]
		for j := 0; j < len(last)-1; j++ {
			if last[j] == last[j+1] {
				count++
				continue
			}
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(last[j])
			count = 1
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(last[len(last)-1])
		if len(dp)-1 < i {
			dp = append(dp, sb.String())
		}
	}
	return dp[n-1]
}
