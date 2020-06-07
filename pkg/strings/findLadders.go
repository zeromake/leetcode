package strings

// FindLadders 单词接龙 II https://leetcode-cn.com/problems/word-ladder-ii
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	var set = make(map[string]bool)
	for _, str := range wordList {
		set[str] = false
	}
	mp := make(map[string]map[string]bool)

	if _, ok := set[endWord]; !ok {
		return nil
	}
	beginWords := make(map[string]bool)
	endWords := make(map[string]bool)
	beginWords[beginWord] = true
	endWords[endWord] = true

	var ret [][]string
	if !buildTree(true, beginWords, endWords, set, mp) {
		return nil
	}
	DFS(mp, &ret, beginWord, endWord, []string{beginWord})
	return ret
}

func DFS(mp map[string]map[string]bool, ret *[][]string, begins string, end string, cache []string) {
	if sons, ok := mp[begins]; ok {
		for k, _ := range sons {
			cache = append(cache, k)
			if k == end {
				tmp := make([]string, len(cache))
				copy(tmp, cache)
				*ret = append(*ret, tmp)
			}
			DFS(mp, ret, k, end, cache)
			cache = cache[0 : len(cache)-1]
		}
	}
}

func buildTree(is bool, beginWords map[string]bool, endWords map[string]bool,
	wordList map[string]bool, mp map[string]map[string]bool) bool {

	if len(beginWords) == 0 {
		return false
	}
	ret := false

	for str, _ := range beginWords {
		delete(wordList, str)
	}

	var next = map[string]bool{}

	for str, _ := range beginWords {
		var sons map[string]bool
		bytes := []byte(str)
		for index, _ := range bytes {
			old := bytes[index]
			for i := 'a'; i <= 'z'; i++ {
				bytes[index] = byte(i)
				s := string(bytes)

				//还有路能走 or 牵桥成功 的话，进行连线
				if _, ok := wordList[s]; ok {
					//牵桥成功
					if _, ok := endWords[s]; ok {
						ret = true
					}
					next[s] = true
					var root, add string
					if is {
						root = str
						add = s
					} else {
						root = s
						add = str
					}
					if _, ok := mp[root]; ok {
						sons = mp[root]
					} else {
						sons = make(map[string]bool)
						mp[root] = sons
					}
					sons[add] = true
				}
				bytes[index] = old
			}
		}

	}
	if ret {
		return ret
	}
	if len(next) < len(endWords) {
		return buildTree(is, next, endWords, wordList, mp)
	} else {
		return buildTree(!is, endWords, next, wordList, mp)
	}
}
