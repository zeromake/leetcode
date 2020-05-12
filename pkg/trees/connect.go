package trees

import (
	"container/list"
	"strconv"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n *Node) String() string {
	sb := &strings.Builder{}
	sb.WriteByte('[')
	node := n
	for node != nil {
		next := node
		for next != nil {
			sb.WriteString(strconv.FormatInt(int64(next.Val), 10))
			sb.WriteByte(',')
			next = next.Next
		}
		if node.Left != nil || node.Right != nil {
			sb.WriteString("#,")
			if node.Left != nil {
				node = node.Left
			} else if node.Right != nil {
				node = node.Right
			}
		} else {
			sb.WriteByte('#')
			node = nil
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

// Connect 填充每个节点的下一个右侧节点指针 https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/ , https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
func Connect(root *Node) *Node {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		lens := queue.Len() - 1
		for i := 0; i <= lens; i++ {
			cur := queue.Front()
			queue.Remove(cur)
			node := cur.Value.(*Node)
			if i < lens {
				node.Next = queue.Front().Value.(*Node)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

func Connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	var (
		most = root
		prev *Node
		curr = root
	)
	processChild := func(childNode *Node) {
		if childNode != nil {
			if prev != nil {
				prev.Next = childNode
			} else {
				most = childNode
			}
			prev = childNode
		}
	}
	for most != nil {
		prev = nil
		curr = most
		most = nil
		for curr != nil {
			processChild(curr.Left)
			processChild(curr.Right)
			curr = curr.Next
		}
	}
	return root
}
