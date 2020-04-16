package trees

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type TreeNodeList []*TreeNode

func (t TreeNodeList) String() string {
	sb := &strings.Builder{}
	for _, rr := range t {
		sb.WriteString("[")
		sb.WriteString(rr.String())
		sb.WriteString("],")
	}
	return sb.String()
}

func (node *TreeNode) String() string {
	if node == nil {
		return "null"
	}
	sb := &strings.Builder{}
	sb.WriteString(strconv.FormatInt(int64(node.Val), 10))
	if node.Right != nil || node.Left != nil {
		sb.WriteString(",")
		sb.WriteString(node.Left.String())
		sb.WriteString(",")
		sb.WriteString(node.Right.String())
	}
	return sb.String()
}

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal
func InorderTraversal(root *TreeNode) []int {
	var result []int
	stack := list.New()
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushFront(root)
			root = root.Left
		}
		e := stack.Front()
		stack.Remove(e)
		root = e.Value.(*TreeNode)
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
