package arrays

import "sort"

// CombinationSum 组合总和 https://leetcode-cn.com/problems/combination-sum-ii
func CombinationSum(nums []int, target int) [][]int {
	var result [][]int
	var lens = len(nums)
	sort.Ints(nums)
	dfs(nums, nil, 0, lens, target, &result)
	return result
}

func dfs(nums []int, path []int, left, lens, residue int, result *[][]int)  {
	if residue == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := left; i < lens; i++ {
		if residue < nums[i] {
			break
		}
		dfs(nums, append(path, nums[i]), i, lens, residue - nums[i], result)
	}
}
