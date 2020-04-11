package arrays

// Combine 组合 https://leetcode-cn.com/problems/combinations
func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	subsetsDfs(&result, nil, 0, nums, len(nums))
	return result
}

func subsetsDfs(result *[][]int, list []int, start int, nums []int, k int) {
	temp := make([]int, len(list))
	copy(temp, list)
	*result = append(*result, temp)
	if len(list) == k {
		return
	}
	for i := start; i < k; i++ {
		num := nums[i]
		subsetsDfs(result, append(list, num), i+1, nums, k)
	}
}
