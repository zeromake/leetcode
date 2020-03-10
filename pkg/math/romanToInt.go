package math

// 罗马数字转数字 https://leetcode-cn.com/problems/roman-to-integer/
func RomanToInt(s string) int {
	romans := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	num := 0
	pre := romans[s[0]]
	for i := 1; i < len(s); i++ {
		v := romans[s[i]]
		if pre < v {
			num -= pre
		} else {
			num += pre
		}
		pre = v
	}
	num += pre
	return num
}
