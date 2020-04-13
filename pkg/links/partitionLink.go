package links

// Partition 分割链表 https://leetcode-cn.com/problems/partition-list
func PartitionLink(head *ListNode, x int) *ListNode {
	after := &ListNode{}
	before := &ListNode{}
	afterRoot := after
	beforeRoot := before
	for head != nil {
		if head.Val >= x {
			after.Next = head
			after = after.Next
		} else {
			before.Next = head
			before = before.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterRoot.Next
	return beforeRoot.Next
}
