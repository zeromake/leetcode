package trees

// IsSymmetric 对称二叉树 https://leetcode-cn.com/problems/symmetric-tree
func IsSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
}
