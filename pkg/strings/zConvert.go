package strings

import "strings"

// ZConvert Z 字形变换 https://leetcode-cn.com/problems/zigzag-conversion/
func ZConvert(s string, rows int) string {
	if rows == 1 {
		return s
	}
	var (
		result = make([]string, rows)
		flag = -1
		row = 0
	)
	for i := 0; i < len(s); i++ {
		c := s[i]
		result[row] += string(c)
		// 实际上就是 0 为正续写入，rows - 1 为倒序写入长度为 rows 的结果
		if row == 0 || row == rows - 1 {
			flag = -flag
		}
		row += flag
	}
	return strings.Join(result, "")
}
