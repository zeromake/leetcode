package links

// RotateRight 旋转链表 https://leetcode-cn.com/problems/rotate-list
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lens := 1
	cur := head
	for cur.Next != nil {
		lens += 1
		cur = cur.Next
	}
	k %= lens
	root := &ListNode{
		Next: head,
	}
	if k > 0 {
		cur.Next = head
		cur = head
		i := 0
		for cur.Next != nil {
			if i == lens-k-1 {
				break
			}
			cur = cur.Next
			i++
		}
		root.Next = cur.Next
		cur.Next = nil
	}
	return root.Next
}
