package trees

// RecoverTree 恢复二叉搜索树 https://leetcode-cn.com/problems/recover-binary-search-tree
func RecoverTree(root *TreeNode) {
	var (
		x, y, pred, predecessor *TreeNode
	)
	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}
