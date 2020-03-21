package arrays

import "sort"

// CombinationSum2 组合总和2 https://leetcode-cn.com/problems/combination-sum-ii
func CombinationSum2(nums []int, target int) [][]int {
	var result [][]int
	var lens = len(nums)
	sort.Ints(nums)
	dfs2(nums, nil, 0, lens, target, &result)
	return result
}

func dfs2(nums []int, path []int, left, lens, residue int, result *[][]int)  {
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
		if i > left && nums[i] == nums[i - 1] {
			continue
		}
		dfs2(nums, append(path, nums[i]), i + 1, lens, residue - nums[i], result)
	}
}
