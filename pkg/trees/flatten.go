package trees

// Flatten 二叉树展开为链表树 https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenDfs(root)
}

func flattenDfs(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	// 取出左右节点
	right := root.Right
	left := root.Left
	// 记得清理原有节点否则会出现循环
	root.Left = nil
	root.Right = nil
	// 把左边的节点压平挂到 Right
	if left != nil {
		root.Right = flattenDfs(left)
	}
	// 取得最后一个节点
	node := root
	for node.Right != nil {
		node = node.Right
	}
	// 挂载右边的节点
	if right != nil {
		node.Right = flattenDfs(right)
	}
	return root
}
