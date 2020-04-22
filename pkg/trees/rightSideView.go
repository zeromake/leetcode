package trees

import "container/list"

// RightSideView 二叉树右视图 https://leetcode-cn.com/problems/binary-tree-right-side-view
func RightSideView(root *TreeNode) []int {
	var cache = make(map[int]int, 0)
	if root == nil {
		return nil
	}
	nodeQueue := list.New()
	depthQueue := list.New()
	nodeQueue.PushBack(root)
	depthQueue.PushBack(0)
	maxDepth := -1
	for nodeQueue.Len() > 0 {
		e := nodeQueue.Front()
		nodeQueue.Remove(e)
		node := e.Value.(*TreeNode)
		e = depthQueue.Front()
		depthQueue.Remove(e)
		depth := e.Value.(int)
		if depth > maxDepth {
			maxDepth = depth
		}
		cache[depth] = node.Val
		if node.Left != nil {
			nodeQueue.PushBack(node.Left)
			depthQueue.PushBack(depth + 1)
		}
		if node.Right != nil {
			nodeQueue.PushBack(node.Right)
			depthQueue.PushBack(depth + 1)
		}
	}
	var result []int
	for i := 0; i <= maxDepth; i++ {
		result = append(result, cache[i])
	}
	return result
}
