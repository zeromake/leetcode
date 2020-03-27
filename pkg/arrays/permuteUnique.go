package arrays

import "sort"

// PermuteUnique 全排列II https://leetcode-cn.com/problems/permutations-ii
func PermuteUnique(nums []int) [][]int {
	var result [][]int
	lens := len(nums)
	if lens == 0 {
		return result
	}
	used := make([]bool, lens)
	path := make([]int, 0)
	sort.Ints(nums)
	permuteDfs2(nums, path, &result, used, lens, 0)
	return result
}

func permuteDfs2(nums, path []int, result *[][]int, used []bool, lens, depth int) {
	if depth == lens {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	for i := 0; i < lens; i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
			continue
		}
		used[i] = true
		permuteDfs2(nums, append(path, nums[i]), result, used, lens, depth + 1)
		used[i] = false
	}
}
