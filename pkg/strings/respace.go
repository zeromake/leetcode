package strings

import (
	"math"

	"github.com/zeromake/leetcode/pkg/utils"
)

const (
	P    = math.MaxInt32
	BASE = 41
)

// Respace 恢复空格 https://leetcode-cn.com/problems/re-space-lcci/
func Respace(dictionary []string, sentence string) int {
	hashValues := map[int]bool{}
	for _, word := range dictionary {
		hashValues[getHash(word)] = true
	}
	f := make([]int, len(sentence)+1)
	for i := 1; i < len(f); i++ {
		f[i] = len(sentence)
	}
	for i := 1; i <= len(sentence); i++ {
		f[i] = f[i-1] + 1
		hashValue := 0
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1]-'a') + 1
			hashValue = (hashValue*BASE + t) % P
			if hashValues[hashValue] {
				f[i] = utils.MinInt(f[i], f[j-1])
			}
		}
	}
	return f[len(sentence)]
}

func getHash(s string) int {
	hashValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		hashValue = (hashValue*BASE + int(s[i]-'a') + 1) % P
	}
	return hashValue
}
