package trees

import "container/list"

// IsSameTree 相同树 https://leetcode-cn.com/problems/same-tree
func IsSameTree(p, q *TreeNode) bool {
	deq := list.New()
	deq.PushBack([2]*TreeNode{p, q})
	for deq.Len() > 0 {
		e := deq.Front()
		deq.Remove(e)
		ee := e.Value.([2]*TreeNode)
		p, q = ee[0], ee[1]
		if !checkTree(p, q) {
			return false
		}
		if p != nil {
			deq.PushBack([2]*TreeNode{p.Left, q.Left})
			deq.PushBack([2]*TreeNode{p.Right, q.Right})
		}
	}
	return true
}

func checkTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val
}
