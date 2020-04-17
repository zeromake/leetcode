package trees

func helper(node *TreeNode, lower, upper *int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}
	if !helper(node.Right, &val, upper) {
		return false
	}
	if !helper(node.Left, lower, &val) {
		return false
	}
	return true
}

// IsValidBST 验证二叉搜索树 https://leetcode-cn.com/problems/validate-binary-search-tree
func IsValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}
