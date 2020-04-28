package math

// SingleNumbers 数组中只出现一次的两个元素 https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof
func SingleNumbers(nums []int) []int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	div := 1
	for i := 0; i < 32; i++ {
		if (div & ret) != 0 {
			break
		}
		div <<= 1
	}
	first, second := 0, 0
	for i := range nums {
		if (div & nums[i]) != 0 {
			first ^= nums[i]
		} else {
			second ^= nums[i]
		}
	}
	return []int{first, second}
}
