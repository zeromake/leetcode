package strings

//type pair struct {
//	Index int
//	Char byte
//}
//
//func MinWindow(s, t string) string {
//	if len(s) == 0 || len(t) == 0 {
//		return ""
//	}
//	dict := make(map[byte]int, 0)
//	for i := 0; i < len(t); i++ {
//		dict[t[i]]++
//	}
//	required := len(dict)
//	var filtered []*pair
//	for i := 0; i < len(s); i++ {
//		c := s[i]
//		if _, ok := dict[c]; ok {
//			filtered = append(filtered, &pair{i, c})
//		}
//	}
//	left, right, formed := 0, 0, 0
//	counts := make(map[byte]int, 0)
//	result := []int{-1, 0, 0}
//	for right < len(filtered) {
//		c := filtered[right].Char
//		counts[c]++
//		count := counts[c]
//		if v, ok := dict[c]; ok && count == v {
//			formed ++
//		}
//		for left <= right && formed == required {
//			c = filtered[left].Char
//			end := filtered[right].Index
//			start := filtered[left].Index
//			if result[0] == -1 || end - start + 1 < result[0] {
//				result[0] = end - start + 1
//				result[1] = start
//				result[2] = end
//			}
//			counts[c]--
//			if v, ok := dict[c]; ok && counts[c] < v {
//				formed--
//			}
//			left++
//		}
//		right++
//	}
//	if result[0] == -1 {
//		return ""
//	}
//	return s[result[1]: result[2]+1]
//}

// MinWindow 最小覆盖字符串 https://leetcode-cn.com/problems/minimum-window-substring
func MinWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := [256]int{}
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	l, count, max, result := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			result = s[l : r+1]
		}
	}
	return result
}
