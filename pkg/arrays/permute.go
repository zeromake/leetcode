package arrays

// Permute 全排列 https://leetcode-cn.com/problems/permutations
func Permute(nums []int) [][]int {
	var result [][]int
	lens := len(nums)
	if lens == 0 {
		return result
	}
	used := make([]bool, lens)
	path := make([]int, 0)
	permuteDfs(nums, path, &result, used, lens, 0)
	return result
}

func permuteDfs(nums, path []int, result *[][]int, used []bool, lens, depth int) {
	if depth == lens {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	for i := 0; i < lens; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		permuteDfs(nums, append(path, nums[i]), result, used, lens, depth+1)
		used[i] = false
	}
}
