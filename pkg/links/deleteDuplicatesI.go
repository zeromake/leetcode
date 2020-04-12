package links

// DeleteDuplicatesI https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
func DeleteDuplicatesI(head *ListNode) *ListNode {
	root := &ListNode{
		Next: head,
	}
	start := root
	end := root
	for start != nil && start.Next != nil {
		val := start.Next.Val
		end = start.Next
		for end != nil && end.Next != nil {
			endVal := end.Next.Val
			if endVal != val {
				break
			}
			end = end.Next
		}
		if start.Next != end && end != nil {
			start.Next = end
		} else {
			start = start.Next
		}
	}
	return root.Next
}
