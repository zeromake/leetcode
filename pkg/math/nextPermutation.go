package math

// NextPermutation 下一个排列 https://leetcode-cn.com/problems/next-permutation/
func NextPermutation(nums []int) {
	// 根本没懂: https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
	if len(nums) <= 1 {
		return
	}
	lens := len(nums)
	i, j, k := lens - 2, lens - 1, lens - 1
	//
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 从后往前找比到第一个比 i 的数小的进行交换
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 如果上面 i = 0 这里直接是排序效果，对 i 后面到结尾进行倒序
	for n, m := j, len(nums) - 1; n < m; n, m = n + 1, m - 1 {
		nums[n], nums[m] = nums[m], nums[n]
	}
}
