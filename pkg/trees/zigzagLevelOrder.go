package trees

import "container/list"

// ZigzagLevelOrder 二叉树的锯齿性层遍历 https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
func ZigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	var dfs func(node *TreeNode, level int)
	if root == nil {
		return result
	}
	dfs = func(node *TreeNode, level int) {
		if level >= len(result) {
			result = append(result, []int{node.Val})
		} else {
			if level%2 == 0 {
				result[level] = append(result[level], node.Val)
			} else {
				result[level] = append([]int{node.Val}, result[level]...)
			}
		}
		nodes := []*TreeNode{node.Left, node.Right}
		for _, nextNode := range nodes {
			if nextNode != nil {
				dfs(nextNode, level+1)
			}
		}
	}
	dfs(root, 0)
	return result
}

func ZigzagLevelOrder2(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		lens := queue.Len()
		lists := make([]int, lens)
		for i := 0; i < lens; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if level%2 == 0 {
				lists[i] = node.Val
			} else {
				lists[lens-i-1] = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
		result = append(result, lists)
	}
	return result
}
