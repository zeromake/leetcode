package math

func Multiply(num1, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	var (
		len1 = len(num1)
		len2 = len(num2)
		sum  = make([]byte, len1+len2)
	)
	// 123
	// 456
	// 2, 2: 3 * 6 = 18
	// [0, 0, 0, 0, 1, 8]
	// 2, 1: 3 * 5 + 1 = 16
	// [0, 0, 0, 1, 6, 8]
	// 2, 0: 3 * 4 + 1 = 13
	// [0, 0, 1, 3, 6, 8]
	// 1, 2: 2 * 6 + 6 = 18
	// [0, 0, 1, 4, 8, 8]
	for i := len1 - 1; i >= 0; i-- {
		c1 := num1[i] - '0'
		for j := len2 - 1; j >= 0; j-- {
			c2 := num2[j] - '0'
			num := c1*c2 + sum[i+j+1]
			sum[i+j+1] = num % 10
			sum[i+j] += num / 10
		}
	}
	var start = 0
	if len(sum) > 0 && sum[0] == 0 {
		start = 1
	}
	for i := 0; i < len(sum); i++ {
		sum[i] += '0'
	}
	return string(sum[start:])
}
