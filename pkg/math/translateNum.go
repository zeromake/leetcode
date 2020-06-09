package math

import "strconv"

// TranslateNum 把数字翻译成字符串 https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
func TranslateNum(num int) int {
	if num < 10 {
		return 1
	}
	mod := num % 100
	if mod < 26 && mod > 9 {
		return TranslateNum(num/10) + TranslateNum(num/100)
	}
	return TranslateNum(num / 10)
}

func TranslateNum2(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

func TranslateNum3(num int) int {
	p, q, r := 0, 0, 1
	i := 0
	cur := 0
	pre := 0
	for num > 0 {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			i++
			continue
		}
		cur = num % 10
		num /= 10
		tmp := cur*10 + pre
		if tmp <= 25 && tmp >= 10 {
			r += p
		}
		i++
	}
	return r
}
