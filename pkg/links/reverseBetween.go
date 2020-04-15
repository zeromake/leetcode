package links

// ReverseBetween 反转链表区间 https://leetcode-cn.com/problems/reverse-linked-list-ii/
func ReverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return nil
	}
	var (
		cur  = head
		prev *ListNode
	)
	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}
	con, tail := prev, cur
	for n > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		n--
	}
	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = cur
	return head
}
