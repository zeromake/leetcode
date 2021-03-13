package math

// Calculate 基本计算器 https://leetcode-cn.com/problems/basic-calculator/
func Calculate(s string) int {
	var (
		result int
		ops    = []int{1}
		sign   = 1
		n      = len(s)
	)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			// 跳过空格
			i++
		case '+':
			// 加号不进行反转符号
			sign = ops[len(ops)-1]
			i++
		case '-':
			// 减号反转符号
			sign = -ops[len(ops)-1]
			i++
		case '(':
			// 进入括号需要单独设置符号
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			// 读出连续的数字
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			// 根据 sign 做加或减的操作
			result += sign * num
		}
	}
	return result
}
