package trees

// NumTrees 不同的二叉搜索树 https://leetcode-cn.com/problems/unique-binary-search-trees
func NumTrees(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result = result * 2 * (2*i + 1) / (i + 2)
	}
	return result
}
