package strings

import "strings"

func ReverseWords(s string) string {
	sb := make([]string, 0)
	start, end := len(s)-1, len(s)-1
	for ; start >= 0; start-- {
		c := s[start]
		if c != ' ' {
			continue
		}
		if end-start > 0 {
			sb = append(sb, s[start+1:end+1])
		}
		end = start - 1
	}
	if end-start > 0 {
		sb = append(sb, s[start+1:end+1])
	}
	return strings.Join(sb, " ")
}
