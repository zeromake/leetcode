package math

// SingleNumber 只出现一次的数字 https://leetcode-cn.com/problems/single-number
func SingleNumber(nums []int) (result int) {
	for _, num := range nums {
		result ^= num
	}
	return
}
