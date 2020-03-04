package links

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s = ""
	cur := l
	for cur != nil {
		s += fmt.Sprintf("%d", cur.Val)
		cur = cur.Next
	}
	return s
}

// AddTwoNumbers 两数相加 https://leetcode-cn.com/problems/add-two-numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		pre = new(ListNode)
		cur = pre
		mod = 0
		x   int
		y   int
		sum int
	)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		sum = x + y + mod
		mod = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
	}
	if mod > 0 {
		cur.Next = &ListNode{
			Val: mod,
		}
	}
	return pre.Next
}
