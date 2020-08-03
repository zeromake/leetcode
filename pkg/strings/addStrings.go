package strings

// AddStrings 数字字符串相加 https://leetcode-cn.com/problems/add-strings/
func AddStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m < n {
		m, n = n, m
		num1, num2 = num2, num1
	}
	pre := byte(0)
	var result []byte
	var offset = m - n
	for i := m - 1; i >= 0; i-- {
		a := num1[i] - '0'
		b := byte(0)
		if i >= offset {
			b = num2[i-offset] - '0'
		}
		sum := a + b + pre
		result = append(result, (sum%10)+'0')
		pre = sum / 10
	}
	if pre != 0 {
		result = append(result, pre+'0')
	}
	lens := len(result)
	// 对结果倒序
	for i := 0; i < lens/2; i++ {
		result[i], result[lens-1-i] = result[lens-1-i], result[i]
	}
	return string(result)
}
