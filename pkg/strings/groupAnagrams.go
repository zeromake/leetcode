package strings

// GroupAnagrams  https://leetcode-cn.com/problems/group-anagrams
func GroupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string, 0)
	for _, s := range strs {
		bytes := make([]byte, 26)
		for i := 0; i < len(s); i++ {
			bytes[s[i]-'a']++
		}
		key := string(bytes)
		v, ok := ans[key]
		if !ok {
			v = make([]string, 0)
		}
		ans[key] = append(v, s)
	}
	list := make([][]string, 0)
	for _, v := range ans {
		list = append(list, v)
	}
	return list
}
