package math

// isPalindrome 是否为回文数 https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(num int) bool {
	if num == 0 {
		return true
	}
	// 边界处理
	if num < 0 || (num%10) == 0 {
		return false
	}
	var (
		revertedNumber = 0
	)
	// 分解数字并倒序
	for num > revertedNumber {
		revertedNumber = revertedNumber*10 + num%10
		num /= 10
	}
	return num == revertedNumber || num == revertedNumber/10
}
