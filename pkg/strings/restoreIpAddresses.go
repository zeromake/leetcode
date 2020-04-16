package strings

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	var result []string
	if s == "" {
		return result
	}
	back(s, 0, nil, &result)
	return result
}

func back(s string, pos int, cur []string, result *[]string) {
	if len(cur) == 4 {
		if pos == len(s) {
			*result = append(*result, strings.Join(cur, "."))
		}
		return
	}
	for i := 1; i <= 3; i++ {
		if pos+i > len(s) {
			break
		}
		segment := s[pos : pos+i]
		if i > 1 && segment[0] == '0' {
			continue
		}
		if i == 3 {
			ii, _ := strconv.Atoi(segment)
			if ii > 255 {
				continue
			}
		}
		back(s, pos+i, append(cur, segment), result)
	}
}
