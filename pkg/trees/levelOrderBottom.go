package trees

import "container/list"

// LevelOrderBottom 二叉树的层序遍历的层倒序 https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
func LevelOrderBottom(root *TreeNode) [][]int {
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
	right := len(result) - 1
	left := 0
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
