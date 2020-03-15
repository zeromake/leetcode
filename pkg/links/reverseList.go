package links

// ReverseList 倒转链表 https://leetcode-cn.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {
	root := &ListNode{}
	for head != nil {
		next := head.Next
		root.Next, head.Next = head, root.Next
		head = next
	}
	return root.Next
}
