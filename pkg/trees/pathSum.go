package trees

// PathSum 路径总和 https://leetcode-cn.com/problems/path-sum-ii
func PathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	dfs(root, sum, []int{}, &result)
	return result
}

func dfs(root *TreeNode, sum int, stack []int, result *[][]int) {
	if root == nil {
		return
	}

	stack = append(stack, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		*result = append(*result, tmp)
		stack = stack[:len(stack)-1]
		return
	}
	val := sum - root.Val
	dfs(root.Left, val, stack, result)
	dfs(root.Right, val, stack, result)
	stack = stack[:len(stack)-1]
}
