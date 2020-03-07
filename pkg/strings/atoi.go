package strings

import "math"

// Atoi 字符串转数字(容错版) https://leetcode-cn.com/problems/string-to-integer-atoi
func Atoi(str string) int {
	var (
		lens = len(str)
	)
	if lens == 0 {
		return 0
	}
	var (
		nums []byte
		flag = 1
		start bool
	)
	for i := 0; i < lens; i++ {
		b := str[i]
		if b == '+' {
			if start {
				break
			}
			flag = 1
			start = true
		} else if b == '-' {
			if start {
				break
			}
			flag = -1
			start = true
		} else if b >= '0' && b <= '9' {
			start = true
			nums = append(nums, b)
		} else if b == ' ' {
			if start {
				break
			}
			continue
		} else {
			if !start {
				return 0
			}
			break
		}
	}
	lens = len(nums)
	sum := 0
	for _, b := range nums {
		ii := int(b - '0')
		sum = sum * 10 + ii
		// 溢出判断
		if flag == -1 && sum * flag < math.MinInt32 {
			return math.MinInt32
		} else if flag == 1 && sum > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	sum *= flag
	return sum
}
