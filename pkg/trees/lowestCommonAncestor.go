package trees

import (
	"container/list"
	"github.com/zeromake/leetcode/pkg/utils"
)

// LowestCommonAncestor 二叉树最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	queue := list.New()
	cache := list.New()
	if root != nil {
		queue.PushBack(root)
		cache.PushBack(([]*TreeNode)(nil))
	}
	var result = [2][]*TreeNode{nil, nil}
loop:
	for queue.Len() > 0 {
		for i := queue.Len() - 1; i >= 0; i-- {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			e = cache.Front()
			cache.Remove(e)
			l := e.Value.([]*TreeNode)
			l = append(l, node)
			if node.Val == p.Val {
				result[0] = l
			} else if node.Val == q.Val {
				result[1] = l
			}
			if result[0] != nil && result[1] != nil {
				break loop
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
				cache.PushBack(l)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				cache.PushBack(l)
			}
		}
	}
	if result[0] != nil && result[1] != nil {
		for i := utils.MinInt(len(result[0]), len(result[1])) - 1; i >= 0; i-- {
			if result[0][i] == result[1][i] {
				return result[0][i]
			}
		}
	}
	return nil
}
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pPath := make([]*TreeNode, 0)
	qPath := make([]*TreeNode, 0)
	traverse(root, &pPath, p)
	traverse(root, &qPath, q)
	i := 0
	for ; i < len(pPath) && i < len(qPath); i++ {
		if pPath[i].Val != qPath[i].Val {
			break
		}
	}
	return pPath[i-1]
}

func traverse(r *TreeNode, path *[]*TreeNode, v *TreeNode) {
	lPath := len(*path)
	*path = append(*path, r)
	if len(*path) > 0 && (*path)[len(*path)-1].Val == v.Val {
		return
	}
	if r.Left != nil {
		traverse(r.Left, path, v)
	}
	if len(*path) > 0 && (*path)[len(*path)-1].Val == v.Val {
		return
	}
	*path = (*path)[:lPath+1]
	if r.Right != nil {
		traverse(r.Right, path, v)
	}
	if len(*path) > 0 && (*path)[len(*path)-1].Val == v.Val {
		return
	}
	*path = (*path)[:lPath]
}
