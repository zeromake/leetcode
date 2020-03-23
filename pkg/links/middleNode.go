package links

// MiddleNode 获取链表的中间节点 https://leetcode-cn.com/problems/middle-of-the-linked-list
func MiddleNode(head *ListNode) *ListNode {
	ptr := head
	for ptr != nil && ptr.Next != nil {
		ptr = ptr.Next.Next
		head = head.Next
	}
	return head
}
