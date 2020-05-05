package trees

import "container/list"

// MinDepth 最小深度 https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func MinDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		result++
		for i := queue.Len(); i > 0; i-- {
			cur := queue.Front()
			queue.Remove(cur)
			node := cur.Value.(*TreeNode)
			// 终止条件为左右节点都为空
			if node.Left == nil && node.Right == nil {
				return result
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}
