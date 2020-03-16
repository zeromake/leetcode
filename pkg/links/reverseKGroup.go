package links

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	count, root, curr := 0, &ListNode{Next: head}, head
	// 存储上次链表的尾部
	head = root
	// 使用栈
	stack := make([]*ListNode, k)
	for curr != nil {
		// 使用栈存储起来
		count++
		stack[count-1] = curr
		curr = curr.Next
		if count == k {
			// 倒转链表
			for i := k - 1; i > 0; i-- {
				stack[i].Next = stack[i-1]
			}
			// 在链表尾部连接上
			stack[0].Next = curr
			// 重新连接链表头部
			head.Next = stack[k-1]
			// 保证下次的链表倒转能够连接上
			head = stack[0]
			count = 0
		}
	}
	return root.Next
}
