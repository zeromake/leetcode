package arrays

import "sort"

// SubsetsWithDup 子集II https://leetcode-cn.com/problems/subsets-ii/
func SubsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	subsetsWithDupDfs(&result, nil, 0, nums, len(nums))
	return result
}

func subsetsWithDupDfs(result *[][]int, list []int, start int, nums []int, k int) {
	temp := make([]int, len(list))
	copy(temp, list)
	*result = append(*result, temp)
	if len(list) == k {
		return
	}
	for i := start; i < k; i++ {
		num := nums[i]
		if i > start && num == nums[i-1] {
			continue
		}
		subsetsWithDupDfs(result, append(list, num), i+1, nums, k)
	}
}
