package trees

// RecoverFromPreorder 从先序遍历还原二叉树 https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/
func RecoverFromPreorder(S string) *TreeNode {
	var (
		path []*TreeNode
		pos  = 0
	)
	for pos < len(S) {
		level := 0
		for S[pos] == '-' {
			level++
			pos++
		}
		value := 0
		for ; pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos++ {
			value = value*10 + int(S[pos]-'0')
		}
		node := &TreeNode{
			Val: value,
		}
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}
	if len(path) > 0 {
		return path[0]
	}
	return nil
}
