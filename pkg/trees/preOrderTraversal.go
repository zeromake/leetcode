package trees

import "container/list"

// PreOrderTraversal 前序遍历 https://leetcode-cn.com/problems/binary-tree-preorder-traversal
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	var result []int
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return result
}
