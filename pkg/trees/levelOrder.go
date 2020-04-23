package trees

import "container/list"

// LevelOrder 二叉树的层序遍历 https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := list.New()
	var result [][]int
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		level := make([]int, size, size)
		for i := 0; i < size; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			level[i] = node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
