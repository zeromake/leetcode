package links

// HasCycle 环形链表 https://leetcode-cn.com/problems/linked-list-cycle
func HasCycle(head *ListNode) bool {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			return true
		}
	}
	return false
}
