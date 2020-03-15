package links

// MergeTwoLists 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	} else if l2 != nil {
		head.Next = l2
	}
	return root.Next
}
