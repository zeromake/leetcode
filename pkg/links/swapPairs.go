package links

// SwapPairs 两两交换链表中的节点 https://leetcode-cn.com/problems/swap-nodes-in-pairs
func SwapPairs(head *ListNode) *ListNode {
	root := &ListNode{
		Next: head,
	}
	pre := root
	for head != nil && head.Next != nil {
		// 交换
		head.Next, head.Next.Next, pre.Next = head.Next.Next, head, head.Next
		pre, head = head, head.Next
	}
	return root.Next
}
