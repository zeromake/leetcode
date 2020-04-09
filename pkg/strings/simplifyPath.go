package strings

import "strings"

// SimplifyPath 简化路径 https://leetcode-cn.com/problems/simplify-path
func SimplifyPath(path string) string {
	result := make([]string, 0)
	lens := 0
	arr := strings.Split(path, "/")
	for _, s := range arr {
		if s == "." || s == "" {
			continue
		} else if s == ".." {
			if lens > 0 {
				result = result[:lens-1]
				lens--
			}
			continue
		}
		result = append(result, s)
		lens++
	}
	return "/" + strings.Join(result, "/")
}
