package links

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	second := &ListNode{
		Next: head,
	}
	root := second
	for i := 0; first != nil && i < n; i++ {
		first = first.Next
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}
	if second.Next != nil {
		second.Next = second.Next.Next
	}
	return root.Next
}
