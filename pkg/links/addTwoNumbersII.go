package links

import "container/list"

// AddTwoNumbersII 两数相加II https://leetcode-cn.com/problems/add-two-numbers
func AddTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := list.New()
	stack2 := list.New()
	for l1 != nil {
		stack1.PushBack(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2.PushBack(l2.Val)
		l2 = l2.Next
	}
	mod := 0
	var result *ListNode
	for stack1.Len() != 0 || stack2.Len() != 0 || mod != 0 {
		val1, val2 := 0, 0
		if stack1.Back() != nil {
			back := stack1.Back()
			stack1.Remove(back)
			val1 = back.Value.(int)
		}
		if stack2.Back() != nil {
			back := stack2.Back()
			stack2.Remove(back)
			val2 = back.Value.(int)
		}
		cur := val1 + val2 + mod
		mod = cur / 10
		cur %= 10
		node := &ListNode{Val: cur}
		node.Next = result
		result = node
	}
	return result
}
