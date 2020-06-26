package strings

import "github.com/zeromake/leetcode/pkg/utils"

// LadderLength 单词接龙 https://leetcode-cn.com/problems/word-ladder
func LadderLength(beginWord string, endWord string, wordList []string) int {
	var (
		step       = 0
		startQueue = [][]byte{utils.StringToBytes(beginWord)}
		endQueue   = [][]byte{utils.StringToBytes(endWord)}
		startUsed  = make([]bool, len(wordList))
		endUsed    = make([]bool, len(wordList))
		wordMap    = make(map[string]int, 0)
	)
	for i, w := range wordList {
		wordMap[w] = i
	}
	if idx, ok := wordMap[endWord]; !ok {
		return 0
	} else {
		endUsed[idx] = true
	}

	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			chars := make([]byte, len(startQueue[i]))
			copy(chars, startQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]
				var c byte = 'a'
				for ; c <= 'z'; c++ {
					chars[j] = c
					idx, ok := wordMap[string(chars)]
					if !ok || startUsed[idx] {
						continue
					} else {
						if endUsed[idx] {
							return step + 1
						}
						startQueue = append(startQueue, utils.StringToBytes(wordList[idx]))
						startUsed[idx] = true
					}
				}
				chars[j] = o
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}
