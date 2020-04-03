package arrays

// GetPermutation 获取有序全排列的第 k 个排列 https://leetcode-cn.com/problems/permutation-sequence
func GetPermutation(n, k int) string {
	result := make([]byte, n)
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	k -= 1
	// 2
	factors := getFactorial(n - 1)
	for i := n - 1; i >= 0; i-- {
		factor := factors[i]
		result[n-i-1] = byte(nums[k/factor] + '0')
		deleteItem(nums, k/factor)
		k %= factor
	}
	return string(result)
}
func deleteItem(nums []int, index int) {
	copy(nums[index:], nums[index+1:])
}

func getFactorial(n int) []int {
	result := make([]int, n+1)
	iTmp := 1
	result[0] = 1
	for i := 1; i <= n; i++ {
		iTmp *= i
		result[i] = iTmp
	}
	return result
}
