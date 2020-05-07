package trees

// IsSubtree 是否为子树 https://leetcode-cn.com/problems/subtree-of-another-tree
func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtreeHelper(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}
func isSubtreeHelper(s, p *TreeNode) bool {
	if p == nil && s == nil {
		return true
	}
	if p == nil || s == nil {
		return false
	}
	return s.Val == p.Val && isSubtreeHelper(s.Left, p.Left) && isSubtreeHelper(s.Right, p.Right)
}
