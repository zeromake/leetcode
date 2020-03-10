package math

import "strings"

// 数字转罗马数字 https://leetcode-cn.com/problems/integer-to-roman/
func IntToRoman(num int) string {
	nums := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	romans := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}
	builder := &strings.Builder{}
	for i := 0; i < 13; i++ {
		for num >= nums[i] {
			builder.WriteString(romans[i])
			num -= nums[i]
		}
	}
	return builder.String()
}

func IntToRoman2(num int) string {
	d := [4][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		// 由于限制了最大值所以最大3000
		{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}
