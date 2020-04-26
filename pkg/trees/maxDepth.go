package trees

import "container/list"

func MaxDepth(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		for i := queue.Len(); i > 0; i-- {
			cur := queue.Front()
			queue.Remove(cur)
			node := cur.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result++
	}
	return result
}
