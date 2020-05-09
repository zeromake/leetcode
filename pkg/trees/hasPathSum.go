package trees

import "container/list"

// HasPathSum 路径总和 https://leetcode-cn.com/problems/path-sum
func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	sumQueue := list.New()
	queue.PushBack(root)
	sumQueue.PushBack(0)
	for queue.Len() > 0 {
		for i := queue.Len() - 1; i >= 0; i-- {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			e = sumQueue.Front()
			sumQueue.Remove(e)
			last := e.Value.(int)
			val := last + node.Val
			if val == sum && node.Right == nil && node.Left == nil {
				return true
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
				sumQueue.PushBack(val)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				sumQueue.PushBack(val)
			}
		}
	}
	return false
}
